package models

import (
	// gorm mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func (ca *Cad_Articles) GetArticles(db *DB) {
	db.Take(ca)
	return
}
