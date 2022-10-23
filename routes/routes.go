package routes

import (
	"net/http"
	"todo/database"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func Create(ctx *gin.Context) {
    task := ctx.PostForm("task")
	database.InsertTask(task)
	ctx.Redirect(302, "/")
}
