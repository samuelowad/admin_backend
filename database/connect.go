package database

import (
	"github.com/samuelowad/admin_backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("mysql:mysql@/go_admin"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Database = db

	err = db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
	if err != nil {
		return
	}
}
