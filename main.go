package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "Golang Minimalist Blog"
	indexData.Desc = "Now it's very easy!"
	marshal, _ := json.Marshal(indexData)
	_, _ = w.Write(marshal)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	log.Fatal(server.ListenAndServe())
}
