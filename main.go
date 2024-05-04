package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World!"))
	})
	log.Fatal(server.ListenAndServe())
}
