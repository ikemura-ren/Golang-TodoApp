package main

import(
	"todo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    engine:= gin.Default()
	engine.LoadHTMLGlob("views/*.html")
	engine.Static("/assets", "./assets/css")
    engine.GET("/", routes.Home) 
	engine.POST("/new", routes.Create)
	engine.GET("/detail/:Id", routes.Detail)
	engine.POST("/update/:Id", routes.Update)
	engine.POST("/delete/:Id", routes.Delete)
    engine.Run(":8080")
}