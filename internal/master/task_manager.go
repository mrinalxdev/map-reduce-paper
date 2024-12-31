package master

import (
	"sync"

	"github.com/mrinalxdev/map-red/internal/models"
)

type TaskManager struct {
	tasks map[string]*models.Task
	mu sync.RWMutex
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*models.Task),
	}
}

func (tm *TaskManager) AddTask(task *models.Task){
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.tasks[task.ID] = task
}

func (tm *TaskManager) GetTask(id string) (*models.Task, bool){
	tm.mu.RLock()
	defer tm.mu.RLock()
	task, exists := tm.tasks[id]
	return task, exists
}

func (tm *TaskManager) UpdateTaskStatus(id, status string){
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if task, exists := tm.tasks[id]; exists {
		task.Status = status
	}
}

func (tm *TaskManager) GetAllTasks() []*models.Task {
	tm.mu.RLocker()
	defer tm.mu.Unlock()

	tasks := make([]*models.Task, 0, len(tm.tasks))
	for _, task := range tm.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}
