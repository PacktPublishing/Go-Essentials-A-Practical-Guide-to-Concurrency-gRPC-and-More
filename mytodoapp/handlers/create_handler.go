package handlers

import (
	"net/http"

	"example.com/mytodoapp/models"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context){
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = len(models.Tasks) + 1
	models.Tasks = append(models.Tasks, task)
	c.JSON(http.StatusOK, task)
}