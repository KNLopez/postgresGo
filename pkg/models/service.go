package models

import (
	"gorm.io/gorm"
	//import config
	"github.com/KNLopez/postgresGo/pkg/config"
)

var db *gorm.DB

type Supplier struct {
	gorm.Model
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	PlanID uint    `json:"plan_id"`
}

type Plan struct {
	gorm.Model
	Suppliers []Supplier `gorm:"foreignKey:PlanID"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Supplier{})
	db.AutoMigrate(&Plan{})
}

func (s *Supplier) CreateSupplier() *Supplier {
	db.Create(&s)
	return s
}

func GetAllSuppliers() []Supplier {
	var Suppliers []Supplier
	db.Find(&Suppliers)
	return Suppliers
}

func GetSupplierById(Id int64) (*Supplier, *gorm.DB) {
	var getSupplier Supplier
	db := db.Where("id = ?", Id).Find(&getSupplier)
	return &getSupplier, db
}

func (s *Supplier) UpdateSupplier(Id int64) *Supplier {
	db.Save(&s)
	return s
}

func DeleteSupplier(Id int64) Supplier {
	var Supplier Supplier
	db.Where("id = ?", Id).Delete(&Supplier)
	return Supplier
}
