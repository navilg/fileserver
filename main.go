package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	path := flag.String("path", currentDir, "Path to serve")
	port := flag.String("port", "3000", "Port to listen on")
	flag.Parse()

	if _, err := os.Stat(*path); os.IsNotExist(err) {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Serving files under %v on port %v \n", *path, *port)
	// http.HandleFunc("/", root)
	http.Handle("/", http.FileServer(http.Dir(*path)))
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
