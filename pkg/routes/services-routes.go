package routes

import (
	"github.com/gorilla/mux"
	// import controller
	"github.com/KNLopez/postgresGo/pkg/controllers"
)


var RegisterServiceRoutes = func(router *mux.Router) {
	// add controller to route
	router.HandleFunc("/api/service", controllers.CreateService).Methods("POST")
	router.HandleFunc("/api/service", controllers.GetServices).Methods("GET")
	router.HandleFunc("/api/service/{id}", controllers.GetService).Methods("GET")
	router.HandleFunc("/api/service/{id}", controllers.UpdateService).Methods("PUT")
	router.HandleFunc("/api/service/{id}", controllers.DeleteService).Methods("DELETE")
}