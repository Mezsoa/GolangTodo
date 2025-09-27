package models

type Todo struct {
	ID    int    `json:"id"`    // Unique identifier
	Title string `json:"title"` // Todo description
	Done  bool   `json:"done"`  // Completion status
}

// In-memory storage for todos (temporary, until we add a database)
var Todos = []Todo{}
