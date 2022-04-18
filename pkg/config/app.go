package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//The '@' at the end of the password 'Visafone@123' is important if you aren't specifying a network address
	d, err := gorm.Open("mysql", "root:Visafone@123@/go-book?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
