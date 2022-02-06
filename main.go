package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// func root(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello Navratan")
// }

func main() {
	path := flag.String("path", "/tmp", "Path to serve")
	port := flag.String("port", "3000", "Port to listen on")
	flag.Parse()
	fmt.Printf("Serving files under %v on port %v \n", *path, *port)
	// http.HandleFunc("/", root)
	http.Handle("/", http.FileServer(http.Dir(*path)))
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
