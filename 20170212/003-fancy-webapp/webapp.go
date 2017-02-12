package main

import (
	"log"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"api_response":42}`))
}

func docrootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	http.ServeFile(w, r, "/var/docroot"+r.URL.Path)
}

func main() {
	log.Println("I am a fancy web app, find me on http://ADDR:19501")
	http.HandleFunc("/", docrootHandler)
	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":19501", nil)
}
