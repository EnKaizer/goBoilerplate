package models

import (
	db "../db"
	// gorm mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB : gorm db model
type DB = db.DB

type Cad_Articles struct {
	Id   int
	ZdId int
	Name string
}
