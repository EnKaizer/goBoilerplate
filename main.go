package main

import (
	"./models"
	"github.com/gin-gonic/gin"

	"./db"
)

type Cad_Articles = models.Cad_Articles

func main() {
	r := gin.Default()

	r.GET("/teste", getArticles)
	r.Run()
}

func getArticles(g *gin.Context) {
	db := db.ConnectDB()
	defer db.Close()

	articles := Cad_Articles{}
	articles.GetArticles(db)
	g.JSON(200, articles)
}
