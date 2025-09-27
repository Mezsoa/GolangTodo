package controllers

import (
	"net/http"
	"strconv"

	"TodoList/models"

	"github.com/gin-gonic/gin"
)

// ListTodos handles GET requests to list all todos
func ListTodos(c *gin.Context) {
	c.JSON(http.StatusOK, models.Todos)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo

	// Bind the JSON payload to the newTodo struct
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign an ID (one more than the last, or 1 if empty)
	if len(models.Todos) == 0 {
		newTodo.ID = 1
	} else {
		newTodo.ID = models.Todos[len(models.Todos)-1].ID + 1
	}

	// Add the new todo to the slice
	models.Todos = append(models.Todos, newTodo)

	// Respond with the created todo
	c.JSON(http.StatusCreated, newTodo)
}

func GetTodoByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, task := range models.Todos {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo was not found"})

}
