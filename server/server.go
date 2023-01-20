package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sse-poc/instance"
	"strings"

	"github.com/r3labs/sse/v2"
)

func logHTTPRequest(w http.ResponseWriter, r *http.Request) {
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, r.Body); err != nil {
		fmt.Printf("Error: %v", err)
	}

	log.Printf("Got trigger request with Data: %v", buf.String())

	server := instance.SSEServer()
	server.Publish("messages", &sse.Event{
		Data: []byte(buf.String()),
	})
}

func main() {
	instance.Init()
	server := instance.SSEServer()

	server.CreateStream("messages")

	mux := http.NewServeMux()
	mux.HandleFunc("/events", server.ServeHTTP)
	mux.HandleFunc("/trigger", logHTTPRequest)

	addr := ":8080"
	log.Println("Starting server on", addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}

}
