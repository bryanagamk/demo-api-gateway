package main

import (
	"service1-go/config"
	"service1-go/controllers"
	"service1-go/job"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	jobManager := job.NewJobManager(3, 10)
	controllers.SetJobManager(jobManager)

	// Start workers
	jobManager.StartWorkers()

	r := gin.Default()

	// Routes
	r.GET("/", controllers.Index)
	r.POST("/products", controllers.CreateProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)

	// Run server
	r.Run(":9001")
}
