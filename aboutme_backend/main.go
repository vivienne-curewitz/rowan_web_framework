package main

import (
	"context"
	"crypto/sha256"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed all:frontend/dist
var frontendAssets embed.FS

func getFileSystem() fs.FS {
	public, err := fs.Sub(frontendAssets, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}
	return public
}

// ip tracking stuff here
type IPInfo struct {
	IPAddress string
	Count     int
	Location  IPGeoLocation
}

type IPGeoLocation struct {
	Lat     float64
	Lon     float64
	City    string
	Country string
}

// http://ip-api.com/json/24.48.0.1
func queryIPGeolocation(ip string) (IPGeoLocation, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return IPGeoLocation{}, err
	}
	defer resp.Body.Close()

	var location IPGeoLocation
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return IPGeoLocation{}, err
	}
	return location, nil
}

func insertIPInfo(remoteIP string, IPList map[string]IPInfo) {
	host, _, err := net.SplitHostPort(remoteIP)
	if err == nil {
		remoteIP = host
	}

	log.Printf("Logging IP: %s\n", remoteIP)

	hash := sha256.Sum256([]byte(remoteIP))
	ipKey := fmt.Sprintf("%x", hash)

	ipInfo, exists := IPList[ipKey]
	if exists {
		ipInfo.Count += 1
		IPList[ipKey] = ipInfo
	} else {
		geoInfo, err := queryIPGeolocation(remoteIP)
		if err != nil {
			log.Printf("Failed to query geolocation for IP %s: %v\n", remoteIP, err)
			geoInfo = IPGeoLocation{}
		}
		IPList[ipKey] = IPInfo{
			IPAddress: ipKey,
			Count:     1,
			Location:  geoInfo,
		}
	}
}

func insertIPDataLoop(ctx context.Context, pool *pgxpool.Pool, ipChan chan string) {
	IPList := make(map[string]IPInfo)
	dbWriteTicker := time.NewTicker(5 * time.Second)
	defer dbWriteTicker.Stop()

	for {
		select {
		case remote, ok := <-ipChan:
			if !ok {
				return
			}
			insertIPInfo(remote, IPList)
		case <-dbWriteTicker.C:
			if len(IPList) == 0 {
				continue
			}
			log.Println("Writing DB Info")
			var infos []IPInfo
			for _, info := range IPList {
				infos = append(infos, info)
			}
			if err := StoreIPInfos(ctx, pool, infos); err != nil {
				log.Printf("Failed to store IP infos: %v\n", err)
			}
		case <-ctx.Done():
			return
		}
	}
}

// the only handler we need for now
func getHandler(public fs.FS, remoteIPChan chan string) http.Handler {
	fileServer := http.FileServer(http.FS(public))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Clean the path to prevent directory traversal
		remoteIPChan <- r.RemoteAddr
		p := path.Clean(r.URL.Path)
		if p == "/" {
			p = "index.html"
		} else if len(p) > 0 && p[0] == '/' {
			p = p[1:] // remove leading slash
		}

		// Check if the file exists and is not a directory
		f, err := public.Open(p)
		if err == nil {
			stat, err := f.Stat()
			if err == nil && !stat.IsDir() {
				f.Close()
				fileServer.ServeHTTP(w, r)
				return
			}
			f.Close()
		}

		if _, err := public.Open("index.html"); err == nil {
			r.URL.Path = "/"
			fileServer.ServeHTTP(w, r)
			return
		}

		http.NotFound(w, r)
	})
}

func main() {
	ctx := context.Background()

	// DB Setup
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" {
		dbUser = "devuser"
	}
	if dbPass == "" {
		dbPass = "devpassword"
	}
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbName == "" {
		dbName = "aboutme_db"
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	pool, err := InitDB(ctx, connString)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer pool.Close()

	remoteIPChan := make(chan string)
	public := getFileSystem()
	handler := getHandler(public, remoteIPChan)
	go insertIPDataLoop(ctx, pool, remoteIPChan)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s...", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
