package models

type User struct {
	ID        int    `json:"id"`        // Unique identifier
	FirstName string `json:"firstName"` // user's first name
	LastName  string `json:"lastName"`  // user's last name
	Email     string `json:"email"`     // user's email address
	Password  string `json:"password"`  // user's password (hashed in a real application)
}

// In-memory storage for users (temporary, until we add a database)
var Users = []User{}
