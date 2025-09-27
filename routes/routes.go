package routes

import (
	"TodoList/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the API routes for the application

func SetupRoutes(router *gin.Engine) {
	router.GET("/todos", controllers.ListTodos)
	router.POST("/todos", controllers.CreateTodo)
	router.GET("/todos/:id", controllers.GetTodoByID)
}
