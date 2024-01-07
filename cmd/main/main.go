package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KNLopez/postgresGo/pkg/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Go ORM Tutorial")
	router := mux.NewRouter()
	routes.RegisterSupplierRoutes(router)

	corsOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsOrigins, corsMethods, corsHeaders)(router)))
}
