package data

import (
	model "github.com/go-api-demo/models"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var todos = []model.Todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Sweep Floor", Completed: false},
	{ID: "3", Item: "Type Letter", Completed: false},
	{ID: "4", Item: "Wash Dishes", Completed: true},
}

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func AddTodo(context *gin.Context) {
	var newTodo model.Todo

	// Binds response to newTodo and if there is an error returns
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodoById(id string) (*model.Todo, error) {
	for i, td := range todos {
		if td.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func GetTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := GetTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func ToggleCompleted(context *gin.Context) {
	id := context.Param("id")
	todo, err := GetTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}
