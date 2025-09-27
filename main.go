package main

import (
	"net/http"

	"TodoList/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routes.SetupRoutes(r)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	r.Run(":8080")

}
