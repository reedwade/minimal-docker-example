package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting\n")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/doc_root"))))
}
