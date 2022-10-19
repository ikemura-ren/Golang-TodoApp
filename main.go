package main

import(
	"todo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    engine:= gin.Default()
	engine.LoadHTMLGlob("views/*.html")
	engine.Static("/static", "./static")
    engine.GET("/", routes.Home) 
    engine.Run(":8080")
}