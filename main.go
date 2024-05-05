package main

import (
	"github.com/gorilla/mux"
	"http-service/pkg"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Start server")

	r := mux.NewRouter()
	r.HandleFunc("/api/generate", pkg.Generate).Methods("POST")
	r.HandleFunc("/api/retrieve/{id}", pkg.Retrieve).Methods("GET")
	//http.HandleFunc("/api/generate", pkg.Generate)
	//http.HandleFunc("/api/retrieve", pkg.Retrieve)
	http.Handle("/", r)
	s := http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
