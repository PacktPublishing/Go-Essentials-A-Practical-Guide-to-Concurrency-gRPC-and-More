package routes

import (
	"example.com/mytodoapp/handlers"
	"github.com/gin-gonic/gin"
)

func CreateRoutes(router *gin.Engine){
	router.POST("/tasks", handlers.CreateTask)
}