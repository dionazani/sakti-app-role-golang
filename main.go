package main

import (
	"log"
	"net/http"
	v1 "sakti-app-role/v1/core"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/app-role", v1.InsertAppRole).Methods("POST")
	router.HandleFunc("/api/v1/app-role/{id}", v1.GetAppRoleById).Methods("GET")

	log.Fatal(http.ListenAndServe(":10000", router))
}
