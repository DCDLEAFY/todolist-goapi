package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv" // String Conversion

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	//uuid -> unique id
)

// Type to handle Json data type that will be stored
type todos struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"Done"`
}

// Define temporary in-memory db
var todolist = []todos{
	// {ID: "1", Description: "This is example number 1", Done: false},
	// {ID: "2", Description: "This is example number 2", Done: false},
	// {ID: "3", Description: "This is example number 3", Done: false},
	// {ID: "4", Description: "This is example number 4", Done: false},
}

var counter int = len(todolist) + 1 //Temporary until persistant data exists

func saveDataToFile() error {
	file, _ := json.MarshalIndent(todolist, "", " ")

	error := os.WriteFile("data.json", file, 0660)

	if error != nil {
		log.Default().Println("Writefile command failed!")
		return nil
	}

	return errors.New("failed os write")
}

func SaveData(c *gin.Context) {
	error := saveDataToFile()

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, todolist)
	} else {
		c.IndentedJSON(http.StatusAccepted, todolist)
	}
}

func readFileToData() error {
	file, error := os.ReadFile("data.json")

	if error != nil {
		log.Default().Println("File not found")
		return errors.New("file not found")
	}

	data := &todolist

	_ = json.Unmarshal([]byte(file), &data)

	return nil
}

func ReadData(c *gin.Context) {
	error := readFileToData()

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, todolist)
	} else {
		c.IndentedJSON(http.StatusOK, todolist)
	}

}

// Get Request : Retrieves todo list
func GetList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todolist)
}

// Post Request : Adds todo point to todo list
func AddTodo(c *gin.Context) {
	var newTodo todos
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	newTodo.ID = strconv.Itoa(counter)
	counter++

	todolist = append(todolist, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// Get the todo list by ID
func GetTodo(id string) (*todos, int, error) {
	for i, b := range todolist {
		if b.ID == id {
			return &todolist[i], i, nil
		}
	}
	return nil, 0, errors.New("todo not found")
}

// Get Request : Retrieve todo via Todo ID
func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	todo, _, err := GetTodo(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Todo item not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

// Patch Request : Flips the status of the done json header
func ChangeStateTodo(c *gin.Context) {
	id := c.Param("id")
	todo, _, err := GetTodo(id)

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
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	todo, index, err := GetTodo(id)

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
