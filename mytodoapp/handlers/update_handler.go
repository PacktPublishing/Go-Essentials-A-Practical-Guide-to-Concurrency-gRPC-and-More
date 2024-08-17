package handlers

import (
	"net/http"
	"strconv"

	"example.com/mytodoapp/models"
	"github.com/gin-gonic/gin"
)

func UpdateTask(c *gin.Context){
	taskIDStr := c.Param("id")
	taskID, err := strconv.Atoi(taskIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range models.Tasks {
		if task.ID == taskID {
			models.Tasks[i] = updatedTask
			c.JSON(http.StatusOK, updatedTask)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}