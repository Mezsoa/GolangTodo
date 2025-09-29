package main

import (
	"net/http"

	"TodoList/database"
	"TodoList/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := database.ConnectMongo(); err != nil {
		panic(err)
	}
	// Close on shutdown
	defer database.CloseMongo()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	r.Run(":8080")

}
