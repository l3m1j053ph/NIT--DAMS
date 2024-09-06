package main

import (
    "net/http"
    "log"
)

func main() {
    // Set up routes
    http.HandleFunc("/", serveStaticFiles)
    
    // Start server
    log.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

// Serve static files from the frontend directory
func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        http.ServeFile(w, r, "./frontend/index.html")
        return
    }
    http.ServeFile(w, r, "./frontend/"+r.URL.Path)
}


