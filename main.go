package main

import(
	"todo/routes"
	"todo/database"
	"github.com/gin-gonic/gin"
)

func main() {
    engine:= gin.Default()
	engine.LoadHTMLGlob("views/*.html")
	engine.Static("/static", "./static")
	database.ConnectionDB()
    engine.GET("/", routes.Home) 
    engine.Run(":8080")
}