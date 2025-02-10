// main.go
package main

import (
	"log"
	"taskmanagementnew/Database"
	"taskmanagementnew/Controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	router := gin.Default()

	router.POST("/tasks", controllers.CreateTask)
	router.GET("/tasks/:id", controllers.GetTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.GET("/tasks", controllers.ListTask)

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}