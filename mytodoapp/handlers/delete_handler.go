package handlers

import (
	"net/http"
	"strconv"

	"example.com/mytodoapp/models"
	"github.com/gin-gonic/gin"
)

func DeleteTask(c *gin.Context){
	id := c.Param("id")

	for i, task := range models.Tasks {
		if strconv.Itoa(task.ID) == id {
			models.Tasks = append(models.Tasks[:i],models.Tasks[i+1:]...)
			c.Status(http.StatusOK)
			return
		}
	}

	c.Status(http.StatusNotFound)
}