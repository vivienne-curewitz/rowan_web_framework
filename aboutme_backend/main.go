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

func writeIPInfoToDB() {
	log.Println("Writing DB Info")
	return
}

func insertIPInfo(remoteIP string, IPList map[string]IPInfo) {
	log.Printf("Logging IP: %s\n", remoteIP)
	ipInfo, exists := IPList[remoteIP]
	if exists {
		ipInfo.Count += 1
		IPList[remoteIP] = ipInfo
	} else {
		ipInfo := IPInfo{
			IPAddress: remoteIP,
			Count:     1,
		}
		IPList[remoteIP] = ipInfo
	}
}

func insertIPDataLoop(ipChan chan string) {
	var IPList map[string]IPInfo
	IPList = make(map[string]IPInfo)
	dbWriteTicker := time.NewTicker(5 * time.Second)
	for {
		select {
		case remote, ok := <-ipChan:
			if !ok {
				dbWriteTicker.Stop()
				return
			}
			insertIPInfo(remote, IPList)
		case <-dbWriteTicker.C:
			// write to database
			writeIPInfoToDB()
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
	remoteIPChan := make(chan string)
	public := getFileSystem()
	handler := getHandler(public, remoteIPChan)
	go insertIPDataLoop(remoteIPChan)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s...", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
