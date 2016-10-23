package main

import (
	"log"
	"net/http"
	"../github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public/"))))

	log.Println("Listening at port 8080")
	http.ListenAndServe(":8080", r)
}