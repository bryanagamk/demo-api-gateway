package job

import (
	"log"
	"sync"
)

// Job represents a unit of work
type Job struct {
	ID      int
	Payload interface{}
}

// JobManager manages jobs and workers
type JobManager struct {
	JobQueue   chan Job
	WorkerPool []Worker
	mu         sync.Mutex
	jobCounter int
}

// NewJobManager initializes a new JobManager
func NewJobManager(workerCount int, queueSize int) *JobManager {
	jobQueue := make(chan Job, queueSize)

	manager := &JobManager{
		JobQueue:   jobQueue,
		WorkerPool: make([]Worker, 0, workerCount),
		jobCounter: 0,
	}

	// Initialize workers
	for i := 1; i <= workerCount; i++ {
		worker := Worker{
			ID:       i,
			JobQueue: jobQueue,
		}
		manager.WorkerPool = append(manager.WorkerPool, worker)
	}

	return manager
}

// StartWorkers starts all workers
func (jm *JobManager) StartWorkers() {
	for _, worker := range jm.WorkerPool {
		go worker.Start()
	}
	log.Println("All workers started.")
}

// AddJob adds a new job to the queue
func (jm *JobManager) AddJob(payload interface{}) {
	jm.mu.Lock()
	defer jm.mu.Unlock()

	jm.jobCounter++
	job := Job{
		ID:      jm.jobCounter,
		Payload: payload,
	}

	select {
	case jm.JobQueue <- job:
		log.Printf("Job %d added to queue.\n", job.ID)
	default:
		log.Printf("Job %d dropped. Queue is full.\n", job.ID)
	}
}
