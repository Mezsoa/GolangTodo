package routes

import (
	"TodoList/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the API routes for the application

func SetupRoutes(router *gin.Engine) {

	// TODOS
	router.GET("/todos", controllers.ListTodos)
	router.POST("/todos", controllers.CreateTodo)
	router.GET("/todos/:id", controllers.GetTodoByID)

	// USERS
	router.POST("/newUser", controllers.CreateUser)
	router.GET("/all/users", controllers.ListUsers)
	//	router.GET("/user/:id", controllers.GetUsersByID)
}
