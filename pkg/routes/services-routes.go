package routes

import (
	"github.com/gorilla/mux"
	// import controller
	"github.com/KNLopez/postgresGo/pkg/controllers"
)

var RegisterSupplierRoutes = func(router *mux.Router) {
	// add controller to route
	router.HandleFunc("/api/supplier", controllers.CreateSupplier).Methods("POST")
	router.HandleFunc("/api/supplier", controllers.GetSuppliers).Methods("GET")
	router.HandleFunc("/api/supplier/{id}", controllers.GetSupplier).Methods("GET")
	router.HandleFunc("/api/supplier/{id}", controllers.UpdateSupplier).Methods("PUT")
	router.HandleFunc("/api/supplier/{id}", controllers.DeleteSupplier).Methods("DELETE")
}
