package main

import (
	"fmt"

	"./models"

	"./db"
)

type Cad_Articles = models.Cad_Articles

func main() {
	db := db.ConnectDB()
	defer db.Close()

	articles := Cad_Articles{}
	articles.GetArticles(db)

	fmt.Println(articles, "blah")
}
