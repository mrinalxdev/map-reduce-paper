package master

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/mrinalxdev/map-red/internal/models"
	"github.com/mrinalxdev/map-red/internal/queue"
	"github.com/mrinalxdev/map-red/internal/storage"
)

type Master struct {
	taskManager *TaskManager
	queue *queue.RabbitMQ
	storage *storage.Redis
}

func NewMaster(queue *queue.RabbitMQ, storage *storage.Redis) *Master {
	return &Master{
		// TODO : add logic for NewTaskManager()
		taskManager: NewTaskManager(),
		queue : queue,
		storage : storage,
	}
}

const (
	chunckSize = 1024 * 1024 // these will covert the big data into 1mb chunks
)

func generateID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (m *Master) splitInput(input []byte) [][]byte {
	var chunks [][]byte
	inputLen := len(input)

	for i := 0; i < inputLen; i += chunckSize {
		end := i + chunckSize
		if end > inputLen {
			end = inputLen
		}
		chunks = append(chunks, input[i:end])
	}
	return chunks
}

func (m *Master) monitorTasks(ctx context.Context) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		
		case <- ticker.C:
			tasks := m.taskManager.GetAllTasks()
			mapComplete := true

			for _, task := range tasks {
                if task.Type == models.MapTask && task.Status != "completed" {
                    mapComplete = false
                    break
                }
            }

			if mapComplete {
                if err := m.initiateReducePhase(ctx); err != nil {
                    return fmt.Errorf("reduce phase initiation failed: %v", err)
                }
            }
		}
	}
}

func (m *Master) initiateReducePhase(ctx context.Context) error {
    // Getting all map results from Redis
    results, err := m.storage.GetMapResults(ctx)
    if err != nil {
        return err
    }

    // Group results by key and create reduce tasks
    for key, values := range results {
        task := &models.Task{
            ID:     generateID(),
            Type:   models.ReduceTask,
            Key:    key,
            Values: []string{values}, // In a real implementation, this would be an array of all values for the key
            Status: "pending",
        }

        if err := m.queue.PublishTask("reduce_tasks", task); err != nil {
            return err
        }
        m.taskManager.AddTask(task)
    }

    return nil
}

func (m *Master) Run(ctx context.Context, input []byte) error {
	chunks := m.splitInput(input)
	for _, chunk := range chunks {
		task := &models.Task{
			// TODO : add logic for generateID()
			// done
			ID: generateID(),
			Type : models.MapTask,
			Data : chunk,
			Status : "pending",
		}

		if err := m.queue.PublishTask("map_tasks", task); err != nil {
			return err
		}
		// TODO : add the logic for AddTask function
		// done
		m.taskManager.AddTask(task)
	}

	// TODO : add the logic for monitortasks function
	// done
	return m.monitorTasks(ctx)
}