package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("usage\n pingserver :8080")
		log.Fatal("must provide addr to listen on")
	}
	addr := os.Args[1]
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("index from %s\n", r.RemoteAddr)
		w.Write([]byte("hit index"))
	})
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("ping from %s\n", r.RemoteAddr)
		w.Write([]byte("pong"))
	})
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Printf("listen on %s\n", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to listn %v", err)
	}
}
