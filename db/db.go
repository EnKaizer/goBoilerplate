package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB : struct gorm
type DB struct {
	*gorm.DB
}

// ConnectDB : connecting DB
func ConnectDB() *DB {
	db, err := gorm.Open("mysql", "cms_faq:yxvy3QfT8vjh7HTu@(10.8.5.252)/cms_faq?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}

	return &DB{db}
}
