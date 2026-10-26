package dto

import (
	"time"
)

const (
	StatusToDo Status = "TODO"
	StatusWip  Status = "WIP"
	StatusDone Status = "DO"
)

type Task struct {
	Id               int
	Priority         int
	Name             string
	Discrip          string
	ProgressComplete float64
	Time             time.Time
	Date             time.Time
	Status           Status
}

type Status = string
