package main

import (
            "encoding/json"
	    "net/http"
	)

type Article struct {
  Title string `json:"Title"`
  Desc string `json:"desc"`
  Content string `json:"content"`
}

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
   article := Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"}
   w.Header().Set("Access-Control-Allow-Origin", "*")
   w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
   w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode(article)
}
