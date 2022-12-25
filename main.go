package main

import (
	"github.com/gin-gonic/gin"
	data "github.com/go-api-demo/data"
)

// 1. Set up dependency tracking
// go mod init github.com/go-api-demo

// 2. Download Gin package
// go get github.com/gin-gonic/gin

func main() {
	router := gin.Default()

	router.GET("/todos", data.GetTodos)

	router.GET("/todos/:id", data.GetTodo)

	router.POST("/todos", data.AddTodo)

	router.PATCH("/todos/:id", data.ToggleCompleted)

	router.Run("localhost:9090")
}
