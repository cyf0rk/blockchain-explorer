package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HealthCheck).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "API is up and running")
}
