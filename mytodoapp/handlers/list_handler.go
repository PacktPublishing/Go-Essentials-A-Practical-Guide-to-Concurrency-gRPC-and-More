package handlers

import (
	"net/http"

	"example.com/mytodoapp/models"
	"github.com/gin-gonic/gin"
)

func ListTasks(c *gin.Context){
	c.JSON(http.StatusOK, models.Tasks)
}