package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.Write(w)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting http server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
