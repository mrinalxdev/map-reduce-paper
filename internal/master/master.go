package master

import (
	"context"

	"github.com/mrinalxdev/map-red/internal/models"
	"github.com/mrinalxdev/map-red/internal/queue"
	"github.com/mrinalxdev/map-red/internal/storage"
)

type Master struct {
	taskManager *models.Task
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

func (m *Master) Run(ctx context.Context, input []byte) error {
	chunks := m.splitInput(input)
	for _, chunk := range chunks {
		task := &models.Task{
			// TODO : add logic for generateID()
			ID: generateID(),
			Type : models.MapTask,
			Data : chunk,
			Status : "pending",
		}

		if err := m.queue.PublishTask("map_tasks", task); err != nil {
			return err
		}
		// TODO : add the logic for AddTask function
		m.taskManager.AddTask(task)
	}

	// TODO : add the logic for monitortasks function
	return m.monitorTasks(task)
}