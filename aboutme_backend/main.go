package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"time"
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
}

var IPList map[string]IPInfo

func writeIpInfoToDB() {
	log.Println("Writing DB Info")
	return
}

func insertIpDataLoop(ipChan chan string) {
	dbWriteTicker := time.NewTicker(5 * time.Second)
	for {
		select {
		case remote, ok := <-ipChan:
			if !ok {
				dbWriteTicker.Stop()
				return
			}
			ipInfo, exists := IPList[remote]
			if exists {
				ipInfo.Count += 1
				IPList[remote] = ipInfo
			} else {
				ipInfo := IPInfo{
					IPAddress: remote,
					Count:     1,
				}
				IPList[remote] = ipInfo
			}
		case <-dbWriteTicker.C:
			// write to database
			writeIpInfoToDB()
		}
	}
}

// the only handler we need for now
func getHandler(public fs.FS) http.Handler {
	fileServer := http.FileServer(http.FS(public))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Clean the path to prevent directory traversal
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

		// If file doesn't exist or is a directory, serve index.html (SPA fallback)
		// but only if it doesn't look like a static asset (to avoid weird errors)
		// Or just always serve index.html for SPA as is common.
		// We'll stick to serving index.html but ensure it exists.

		if _, err := public.Open("index.html"); err == nil {
			r.URL.Path = "/"
			fileServer.ServeHTTP(w, r)
			return
		}

		http.NotFound(w, r)
	})
}

func main() {
	public := getFileSystem()
	handler := getHandler(public)
	remoteIPChan := make(chan string)
	go insertIpDataLoop(remoteIPChan)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s...", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
