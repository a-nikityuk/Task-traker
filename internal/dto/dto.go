package dto

import (
	"time"
)

const (
	StatusToDo Status = "todo"
	StatusWip  Status = "in_progress"
	StatusDone Status = "done"
)

type Task struct {
	Id               int
	Name             string
	Priority         int
	Discription      string
	ProgressComplete float64
	Time             time.Time
	Date             time.Time
	Status           Status
}

type Status = string
