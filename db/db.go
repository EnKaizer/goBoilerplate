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
	db, err := gorm.Open("mysql", "root:root@(localhsot)/teste?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}

	return &DB{db}
}
