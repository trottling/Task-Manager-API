package models

import "time"

type Status string

const (
	StatusNew        Status = "new"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          int
	Name        string
	Description string
	Status      Status
	CreatedAt   time.Time
}
