package controllers

import (
	"net/http"
	"service1-go/config"
	"service1-go/job"
	"service1-go/models"

	"github.com/gin-gonic/gin"
)

var JobManager *job.JobManager

func SetJobManager(manager *job.JobManager) {
	JobManager = manager
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Hello, World!"})
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	// Bind JSON input
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Add job to queue
	JobManager.AddJob(product)

	c.JSON(http.StatusOK, gin.H{"message": "Job added to queue"})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.MySql.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.MySql.Save(&product)
	c.JSON(http.StatusOK, product)
}
