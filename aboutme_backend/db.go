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
		count INTEGER DEFAULT 0,
		lat DOUBLE PRECISION,
		lon DOUBLE PRECISION,
		city TEXT,
		country TEXT
	);`
	_, err = pool.Exec(ctx, query)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to create table: %w", err)
	}

	// Ensure columns exist (for existing tables)
	alterQueries := []string{
		"ALTER TABLE ip_info ADD COLUMN IF NOT EXISTS lat DOUBLE PRECISION",
		"ALTER TABLE ip_info ADD COLUMN IF NOT EXISTS lon DOUBLE PRECISION",
		"ALTER TABLE ip_info ADD COLUMN IF NOT EXISTS city TEXT",
		"ALTER TABLE ip_info ADD COLUMN IF NOT EXISTS country TEXT",
	}
	for _, q := range alterQueries {
		if _, err := pool.Exec(ctx, q); err != nil {
			return nil, fmt.Errorf("failed to ensure column exists: %w", err)
		}
	}

	return pool, nil
}

// StoreIPInfos stores a list of IPInfo structs in the database using UPSERT.
func StoreIPInfos(ctx context.Context, pool *pgxpool.Pool, infos []IPInfo) error {
	batch := &pgx.Batch{}
	for _, info := range infos {
		batch.Queue(`
			INSERT INTO ip_info (ip_address, count, lat, lon, city, country)
			VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT (ip_address)
			DO UPDATE SET 
				count = EXCLUDED.count,
				lat = EXCLUDED.lat,
				lon = EXCLUDED.lon,
				city = EXCLUDED.city,
				country = EXCLUDED.country;
		`, info.IPAddress, info.Count, info.Location.Lat, info.Location.Lon, info.Location.City, info.Location.Country)
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
	rows, err := pool.Query(ctx, "SELECT ip_address, count, lat, lon, city, country FROM ip_info")
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var infos []IPInfo
	for rows.Next() {
		var info IPInfo
		if err := rows.Scan(&info.IPAddress, &info.Count, &info.Location.Lat, &info.Location.Lon, &info.Location.City, &info.Location.Country); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		infos = append(infos, info)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return infos, nil
}
