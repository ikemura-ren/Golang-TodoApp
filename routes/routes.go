package routes

import (
	"strconv"
	"todo/database"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	tasklist := database.GetAllTasks()
    ctx.HTML(200, "index.html", gin.H{"tasklist": tasklist})
}


func Create(ctx *gin.Context) {
    task := ctx.PostForm("task")
	database.InsertTask(task)
	ctx.Redirect(302, "/")
}


func Detail(ctx *gin.Context) {
	task_id := ctx.Param("Id")
	id , err := strconv.Atoi(task_id)
		if err != nil {
			panic("ERROR")
		}
    task := database.GetOneTask(id)
	ctx.HTML(200, "detail.html", gin.H{"task": task})
}


func Update(ctx *gin.Context) {
	task_id := ctx.Param("Id")
	id, err := strconv.Atoi(task_id)
        if err != nil {
            panic("ERROR")
        }
	task := ctx.PostForm("task")
	progress := ctx.PostForm("progress")
    database.UpdateTask(id, task, progress)
    ctx.Redirect(302, "/")
	}


func Delete(ctx *gin.Context) {
	task_id := ctx.Param("Id")
	id, err := strconv.Atoi(task_id)
        if err != nil {
            panic("ERROR")
        }
    database.DeleteTask(id)
    ctx.Redirect(302, "/")
}