package routes

import (
	"example.com/mytodoapp/handlers"
	"github.com/gin-gonic/gin"
)

func ListRoutes(router *gin.Engine){
	router.GET("/tasks", handlers.ListTasks)
}