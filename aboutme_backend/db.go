package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// InitDB initializes the database connection and creates the necessary tables.
func InitDB(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Create table if it doesn't exist
	query := `
	CREATE TABLE IF NOT EXISTS ip_info (
		ip_address TEXT PRIMARY KEY,
		count INTEGER DEFAULT 0
	);`
	_, err = pool.Exec(ctx, query)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to create table: %w", err)
	}

	return pool, nil
}

// StoreIPInfos stores a list of IPInfo structs in the database using UPSERT.
func StoreIPInfos(ctx context.Context, pool *pgxpool.Pool, infos []IPInfo) error {
	batch := &pgx.Batch{}
	for _, info := range infos {
		batch.Queue(`
			INSERT INTO ip_info (ip_address, count)
			VALUES ($1, $2)
			ON CONFLICT (ip_address)
			DO UPDATE SET count = EXCLUDED.count;
		`, info.IPAddress, info.Count)
	}

	results := pool.SendBatch(ctx, batch)
	defer results.Close()

	for range infos {
		_, err := results.Exec()
		if err != nil {
			return fmt.Errorf("batch execution failed: %w", err)
		}
	}

	return nil
}

// GetIPInfos retrieves all IPInfo structs from the database.
func GetIPInfos(ctx context.Context, pool *pgxpool.Pool) ([]IPInfo, error) {
	rows, err := pool.Query(ctx, "SELECT ip_address, count FROM ip_info")
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var infos []IPInfo
	for rows.Next() {
		var info IPInfo
		if err := rows.Scan(&info.IPAddress, &info.Count); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		infos = append(infos, info)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return infos, nil
}
