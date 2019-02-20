package common

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"
import "github.com/yowithus/fennekin/models"

var db *gorm.DB
var err error

func InitDB() {
	db, err = gorm.Open("mysql", "yonatan:123456@/fennekin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	migrate()
}

func GetDB() *gorm.DB {
	return db
}

func migrate() {
	db.AutoMigrate(&models.Product{})
}
