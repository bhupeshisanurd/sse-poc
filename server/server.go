package main

import (
	"fmt"
	"io"
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
	method := r.Method

	logMsg := fmt.Sprintf("Method: %v, Body: %v", method, buf.String())
	fmt.Println(logMsg)

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
	mux.HandleFunc("/push", logHTTPRequest)

	http.ListenAndServe(":8080", mux)

}
