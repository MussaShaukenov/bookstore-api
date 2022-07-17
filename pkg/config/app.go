package config

import (
	_ "github/com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	db, err := gorm.Open("simplerest", "root:mussa2002@tcp(127.0.0.1:3306)/example?charset=utf8 &parseTime=True&loc=Local")
	if err != nil { // if there is an error
		panic(err)
	}
	defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}
