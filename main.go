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
	engine.POST("/new", routes.Create)
	engine.POST("/delete/:Id", routes.Delete)
    engine.Run(":8080")
}