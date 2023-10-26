package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Connection made to %s with protocol %s\n", r.Host, r.Proto)
	w.Write([]byte("Hello from Go Server!"))
}

func main() {
	http.HandleFunc("/", handler)

	// HTTP 1.1 Server
	go func() {
		log.Println("HTTP/1.1 server listening on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("HTTP/1.1 server error: %v", err)
		}
	}()

	// Unencrypted HTTP/2 Server with h2c
	go func() {
		h2s := &http2.Server{}
		h1s := &http.Server{
			Addr:    ":8081",
			Handler: h2c.NewHandler(http.HandlerFunc(handler), h2s),
		}
		log.Println("HTTP/2 h2c server listening on :8081")
		if err := h1s.ListenAndServe(); err != nil {
			log.Fatalf("HTTP/2 h2c server error: %v", err)
		}
	}()

	// Block main routine forever
	select {}
}
