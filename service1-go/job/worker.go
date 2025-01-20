package job

import (
	"log"
	"service1-go/config"
	"service1-go/models"
	"time"
)

// Worker processes jobs from the queue
type Worker struct {
	ID       int
	JobQueue chan Job
}

// Start begins processing jobs
func (w *Worker) Start() {
	log.Printf("Worker %d started.\n", w.ID)
	for job := range w.JobQueue {
		log.Printf("Worker %d processing job %d with payload: %+v\n", w.ID, job.ID, job.Payload)

		// Assert job payload to models.Product
		product, ok := job.Payload.(models.Product)
		if !ok {
			log.Printf("Worker %d failed to process job %d: invalid payload type\n", w.ID, job.ID)
			continue
		}

		// Insert product into database
		if err := config.MySql.Create(&product).Error; err != nil {
			log.Printf("Worker %d failed to insert product: %v\n", w.ID, err)
			continue
		}

		log.Printf("Worker %d successfully inserted product: %s\n", w.ID, product.Name)
		time.Sleep(2 * time.Second) // Simulate processing time
		log.Printf("Worker %d completed job %d\n", w.ID, job.ID)
	}
}
