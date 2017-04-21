package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.Write(w)
	fmt.Fprintf(w, "RemoteAddr: %s", r.RemoteAddr)
}

func main() {
	port := flag.Int("port", 8080, "Listen on port")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Println("Starting http server on port:", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
