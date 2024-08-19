package routes

import (
	"example.com/mytodoapp/handlers"
	"github.com/gin-gonic/gin"
)

func UpdateRoutes(router *gin.Engine){
	router.PUT("/tasks/:id", handlers.UpdateTask)
}