package routes

import (
	"net/http"
    "github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "index.html", gin.H{})
}