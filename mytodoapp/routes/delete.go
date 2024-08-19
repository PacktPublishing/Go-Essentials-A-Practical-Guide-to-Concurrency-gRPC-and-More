package routes

import (
	"example.com/mytodoapp/handlers"
	"github.com/gin-gonic/gin"
)

func DeleteRoutes(router *gin.Engine){
	router.DELETE("/tasks/:id", handlers.DeleteTask)
}