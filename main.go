package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

// Type to handle Json data type that will be stored
type todos struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"Done"`
}

// Define temporary in-memory db
var todolist = []todos{
	{ID: "1", Description: "This is example number 1", Done: false},
	{ID: "2", Description: "This is example number 2", Done: false},
	{ID: "3", Description: "This is example number 3", Done: false},
	{ID: "4", Description: "This is example number 4", Done: false},
}

// Get Request : Retrieves todo list
func getList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todolist)
}

// Post Request : Adds todo point to todo list
func addTodo(c *gin.Context) {
	var newTodo todos
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todolist = append(todolist, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// Get the todo list by ID
func getTodo(id string) (*todos, int, error) {
	for i, b := range todolist {
		if b.ID == id {
			return &todolist[i], i, nil
		}
	}
	return nil, 0, errors.New("todo not found")
}

// Get Request : Retrieve todo via Todo ID
func getTodoById(c *gin.Context) {
	id := c.Param("id")
	todo, _, err := getTodo(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Todo item not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

// Patch Request : Flips the status of the done json header
func changeStateTodo(c *gin.Context) {
	id := c.Param("id")
	todo, _, err := getTodo(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Todo item not found"})
		return
	}

	todo.Done = true
	c.IndentedJSON(http.StatusAccepted, todo)
}

// Remove an index from the todos slice
func removeIndex(s []todos, i int) []todos {
	return append(s[:i], s[i+1:]...)
}

// Delete Request : Delete todo from the todo list via ID, only deletes if the json done flag is set to true
func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	todo, index, err := getTodo(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Todo item not found"})
		return
	}

	if todo.Done {
		todolist = removeIndex(todolist, index)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Todo item not done"})
	}
}

func main() {
	router := gin.Default()
	router.GET("/todo/getlist", getList)
	router.GET("/todo/gettodo/:id", getTodoById)
	router.PATCH("/todo/fliptodo/:id", changeStateTodo)
	router.POST("/todo/addtodo", addTodo)
	router.DELETE("todo/deletetodo/:id", deleteTodo)
	router.Run("localhost:8080")
}
