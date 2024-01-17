package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KNLopez/postgresGo/pkg/suppliers/models"
	"github.com/KNLopez/postgresGo/pkg/suppliers/utils"
	"github.com/gorilla/mux"
)

var NewSupplier models.Supplier

func CreateSupplier(w http.ResponseWriter, r *http.Request) {
	CreateSupplier := &models.Supplier{}
	err := utils.ParseBody(r, CreateSupplier)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	NewSupplier := CreateSupplier.CreateSupplier()
	utils.ToJson(w, NewSupplier, http.StatusCreated)
}

func GetSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers := models.GetAllSuppliers()
	utils.ToJson(w, suppliers, http.StatusOK)
}

func GetSupplier(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	supplier, _ := models.GetSupplierById(Id)
	utils.ToJson(w, supplier, http.StatusOK)
}

func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	UpdatedSupplier := &models.Supplier{}
	err = json.NewDecoder(r.Body).Decode(UpdatedSupplier)
	if err != nil {
		fmt.Println("Error while decoding")
	}
	UpdatedSupplier.ID = uint(Id)
	supplier := UpdatedSupplier.UpdateSupplier(Id)
	utils.ToJson(w, supplier, http.StatusOK)
}

func DeleteSupplier(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	Id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	supplier := models.DeleteSupplier(Id)
	utils.ToJson(w, supplier, http.StatusOK)
}
