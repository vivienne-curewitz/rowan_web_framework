package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s...", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
