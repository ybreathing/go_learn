package common

import (
	"fmt"
	"zzy/go-learn/module"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "go_learn"
	username := "root"
	password := "123456"
	charset := "utf8"

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	db, err := gorm.Open(driverName, args)

	if err != nil {
		panic("failed to connect datebase, err:" + err.Error())
	}
	db.AutoMigrate(&module.User{})
	DB = db
	return db

}

func GetDB() *gorm.DB {
	return DB

}
