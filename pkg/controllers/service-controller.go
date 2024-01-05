package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KNLopez/postgresGo/pkg/models"
	"github.com/KNLopez/postgresGo/pkg/utils"
	"github.com/gorilla/mux"
)

var NewService models.Service

func CreateService(w http.ResponseWriter, r *http.Request) {
	CreateService := &models.Service{}
	err := utils.ParseBody(r, CreateService)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	NewService := CreateService.CreateService()
	utils.ToJson(w, NewService, http.StatusCreated)
}

func GetServices(w http.ResponseWriter, r *http.Request) {
	services := models.GetAllServices()
	utils.ToJson(w, services, http.StatusOK)
}

func GetService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	service, _ := models.GetServiceById(Id)
	utils.ToJson(w, service, http.StatusOK)
}

func UpdateService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	UpdateService := &models.Service{}
	err = json.NewDecoder(r.Body).Decode(UpdateService)
	if err != nil {
		fmt.Println("Error while decoding")
	}
	service := UpdateService.UpdateService(Id)
	utils.ToJson(w, service, http.StatusOK)
}

func DeleteService(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	Id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	service := models.DeleteService(Id)
	utils.ToJson(w, service, http.StatusOK)
}


