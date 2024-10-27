package task

import "time"

type Task struct {
	ID int
	Title string
	Description string
	CreatedAt time.Time
}