package database

import (

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connection := "root:root@(127.0.0.1:5432)/rootdb?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(connection))
	if err != nil {
		log.Panic("Error in connection with DataBase: ", err)
	}
	log.Info("Conecction with database begin...")
}
