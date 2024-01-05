package models

import (
	"gorm.io/gorm"
	//import config
	"github.com/KNLopez/postgresGo/pkg/config"
)

var db *gorm.DB

type Service struct {
	gorm.Model
	Name  string `json:"name"`
	Price float64 `json:"price"`
}

type Plan struct {
	gorm.Model
	Services []Service `gorm:"foreignKey:PlanID"`
}


func init() {
	config.Connect()	
	db = config.GetDB()
	db.AutoMigrate(&Service{})
	db.AutoMigrate(&Plan{})
}

func (s *Service) CreateService() *Service {
	db.Create(&s)
	return s
}

func GetAllServices() []Service {
	var services []Service
	db.Find(&services)
	return services
}

func GetServiceById(Id int64) (*Service, *gorm.DB) {
	var getService Service
	db := db.Where("id = ?", Id).Find(&getService)
	return &getService, db
}

func (s *Service) UpdateService(Id int64) *Service {
	db.Save(&s)
	return s
}

func DeleteService(Id int64) Service {
	var service Service
	db.Where("id = ?", Id).Delete(&service)
	return service
}




