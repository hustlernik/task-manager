package main

import (
	"fmt"
	"task-manager/task"
	"sync"
)

func main() {
	tm := task.NewTaskManager()
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		err := tm.AddTask("Learn Go","Complete Go basics and concurrency\n")
		if err != nil {
			fmt.Println("Error adding task:",err)
		}
	}()

	go func() {
		defer wg.Done()
		err := tm.AddTask("Build Project", "Develop a project using Go\n")
		if err !=nil {
			fmt.Println("Error adding task:",err)
		}
	}()

	wg.Wait()

	fmt.Println("Current Tasks:")
	for _, task := range tm.ListTasks() {
		fmt.Printf("ID:%d,Title:%s,Description:%s\n",task.ID,task.Title,task.Description)
	}

	err := tm.DeleteTask(1)
	if err != nil {
		fmt.Println("Error deleting task", err)
	} else {
		fmt.Println("Task deleted successfully.")
	}

	 fmt.Println("Tasks After Deletion:")
    for _, task := range tm.ListTasks() {
        fmt.Printf("ID: %d, Title: %s, Description: %s\n", task.ID, task.Title, task.Description)
    }

}