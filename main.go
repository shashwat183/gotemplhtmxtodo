package main

import (
	"gohtmxtodolist/data"
	"gohtmxtodolist/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	data.Load()
	r := gin.Default()

	// routes
	r.GET("/", handlers.RootGet)
	r.POST("/tasks", handlers.CreateTask)
	r.POST("/tasks/:taskId/toggle", handlers.TaskToggle)
	r.DELETE("/tasks/:taskId", handlers.DeleteTask)
	r.POST("/tasks/clear", handlers.ClearDoneTasks)

	r.Run()
}
