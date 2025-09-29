package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"` // Mongo ObjectID
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"` // NOTE: should be hashed
}

// In-memory storage for users (temporary, until we add a database)
// var Users = []User{}
