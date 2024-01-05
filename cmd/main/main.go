package main

import (
	"log"
	"net/http"

	"github.com/KNLopez/postgresGo/pkg/routes"
	"github.com/gorilla/mux"
)


func main() {
	
	router := mux.NewRouter()
	routes.RegisterServiceRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
	
}