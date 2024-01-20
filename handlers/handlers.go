package handlers

import (
	"gohtmxtodolist/data"
	"gohtmxtodolist/templates"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NewTaskForm struct {
	Title       string `form:"task"`
	Description string `form:"description"`
}

func RootGet(ctx *gin.Context) {
	templates.Root(data.Tasks).Render(ctx, ctx.Writer)
}

func CreateTask(ctx *gin.Context) {
	var f NewTaskForm
	ctx.Bind(&f)
	if f.Title == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}
	id := uuid.NewString()
	task := data.NewTask(f.Title, f.Description)
	data.Tasks.AddTask(id, task)
	templates.Task(id, *task).Render(ctx, ctx.Writer)
}

func TaskToggle(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	task, ok := data.Tasks.TasksMap[taskId]
	if !ok {
		ctx.Status(http.StatusNotFound)
		return
	}
	task.ToggleStatus()
	templates.Task(taskId, *task).Render(ctx, ctx.Writer)
}

func DeleteTask(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	err := data.Tasks.DeleteTask(taskId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}

func ClearDoneTasks(ctx *gin.Context) {
	for id, task := range data.Tasks.TasksMap {
		if task.Status == data.DoneStatus {
			err := data.Tasks.DeleteTask(id)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
	}
	ctx.Status(http.StatusOK)
	templates.TasksContainer(data.Tasks).Render(ctx, ctx.Writer)
}
