package task

type TaskManagerInterface interface {
	 AddTask(title,description string) error
	 ListTasks() []Task
	 DeleteTask(id int) error
}