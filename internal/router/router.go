package router

import (
	"taskmanager/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/tasks", controller.GetTasks)
	r.POST("/tasks", controller.CreateTask)
	r.GET("/tasks/:id", controller.GetTaskByID)
	r.PUT("/tasks/:id", controller.UpdateTask)
	r.DELETE("/tasks/:id", controller.DeleteTask)
}
