package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Connection made to %s with protocol %s\n", r.Host, r.Proto)
	w.Write([]byte("Hello Flyerz!"))
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

	// Unencrypted HTTP/2 Server
	go func() {
		server := &http.Server{
			Addr:    ":8081",
			Handler: http.HandlerFunc(handler),
		}
		http2.ConfigureServer(server, &http2.Server{})
		log.Println("HTTP/2 unencrypted server listening on :8081")
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("HTTP/2 unencrypted server error: %v", err)
		}
	}()

	// // Encrypted HTTP/2 Server
	// go func() {
	// 	server := &http.Server{
	// 		Addr:    ":8082",
	// 		Handler: http.HandlerFunc(handler),
	// 	}
	// 	http2.ConfigureServer(server, &http2.Server{})
	// 	log.Println("HTTP/2 encrypted server listening on :8082")
	// 	if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
	// 		log.Fatalf("HTTP/2 encrypted server error: %v", err)
	// 	}
	// }()

	// Block main routine forever
	select {}
}
