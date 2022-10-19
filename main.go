package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
    engine:= gin.Default()
	engine.LoadHTMLGlob("views/*.html")
	engine.Static("/static", "./static")
    engine.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{})
        })
    engine.Run(":8080")
}