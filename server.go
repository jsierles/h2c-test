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
	w.Write([]byte("Hello Flyerz!"))
}
func main() {
	h2s := &http2.Server{}
	h1s := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(http.HandlerFunc(handler), h2s),
	}
	log.Println("HTTP/2 h2c server listening on :8080")
	if err := h1s.ListenAndServe(); err != nil {
		log.Fatalf("HTTP/2 h2c server error: %v", err)
	}
}
