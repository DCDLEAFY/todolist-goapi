package main

import (
	"github.com/dcdleafy/todolist-goapi/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/todo/getlist", controllers.GetList)
	router.GET("/todo/gettodo/:id", controllers.GetTodoById)
	router.PATCH("/todo/fliptodo/:id", controllers.ChangeStateTodo)
	router.POST("/todo/addtodo", controllers.AddTodo)
	router.DELETE("todo/deletetodo/:id", controllers.DeleteTodo)
	router.PATCH("/todo/savedata", controllers.SaveData)
	router.PATCH("/todo/readdata", controllers.ReadData)
	router.Run("localhost:8080")
}
