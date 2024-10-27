package task

import (
	"fmt"
	"sync"
	"time"
)

type TaskManager struct {
	tasks map[int]Task
	mu sync.Mutex
	nextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:make(map[int]Task),
		nextID:1,
	}
}

func (tm *TaskManager) AddTask(title,description string) error {
    tm.mu.Lock()
		defer tm.mu.Unlock();

		task := Task{
			ID: tm.nextID,
			Title: title,
			Description: description,
			CreatedAt: time.Now(),
		}
		tm.tasks[tm.nextID] = task
		tm.nextID++
		return nil

}

func (tm *TaskManager) ListTasks() []Task {
	tm.mu.Lock();
	defer tm.mu.Unlock()

	tasks := make([]Task,0,len(tm.tasks))
	for _, task := range tm.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (tm *TaskManager) DeleteTask(id int) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if _, exists := tm.tasks[id]; !exists {
		return fmt.Errorf("task with ID %d not found", id)
  }

	delete(tm.tasks,id)
	return nil

}
