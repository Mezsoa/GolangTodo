package controllers

import (
	"context"
	"net/http"
	"time"

	"TodoList/database"
	"TodoList/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(c *gin.Context) {
	var payload models.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// basic validation
	if payload.Email == "" || payload.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is required"})
		return
	}

	// check if email exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := database.UsersColl.FindOne(ctx, bson.D{{"email", payload.Email}}).Err()
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		return
	}
	if err != nil && err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// Insert user (password shoould be hashed in a real app)
	res, err := database.UsersColl.InsertOne(ctx, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create a user"})
		return
	}

	payload.ID = res.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusOK, payload)

}

func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, database.UsersColl)
}

//func GetUsersByID(c *gin.Context) {
//	idString := c.Param("id")
//	id, err := strconv.Atoi(idString)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	for _, user := range database.UsersColl.Find("id") {
//
//	}
//
//	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//  }
