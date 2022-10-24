package routes

import (
	"todo/database"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	tasklist := database.GetAllTasks()
	fmt.Println(tasklist)
    ctx.HTML(200, "index.html", gin.H{"tasklist": tasklist})
}

func Create(ctx *gin.Context) {
    task := ctx.PostForm("task")
	database.InsertTask(task)
	tasklist := database.GetAllTasks()
	ctx.HTML(302, "/", gin.H{"tasklist": tasklist})
}
