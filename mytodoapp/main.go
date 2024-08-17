package main

import (
	"example.com/mytodoapp/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	routes.CreateRoutes(router)
	routes.UpdateRoutes(router)
	routes.DeleteRoutes(router)
	routes.ListRoutes(router)

	router.Run(":8080")
}