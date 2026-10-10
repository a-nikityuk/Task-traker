package main

import (
	"fmt"
	"slices"
	"time"
)

type Task struct {
	_id               int
	_priority         int
	_name             string
	_discrip          string
	_additInfo        string
	_progressComplete float64
	_time             time.Time
	_date             time.Time
	_isDo             bool
}

type Storage interface {
	CreateTask()
	GetTask()
	UpdateTask()
	DeleteTask()
}

type MemoryStorage struct {
	tasks []Task
}

func (ms *MemoryStorage) CreateTask(id, priority int, name, discrip, additInfo string, pComplete float64, time, date time.Time, is_Do bool) {
	task := Task{
		_id:               id,
		_priority:         priority,
		_name:             name,
		_discrip:          discrip,
		_additInfo:        additInfo,
		_progressComplete: pComplete,
		_time:             time,
		_date:             date,
		_isDo:             is_Do,
	}
	ms.tasks = append(ms.tasks, task)
}

func (ms *MemoryStorage) GetTask(id int) Task {
	getTask := ms.tasks[id-1]
	return getTask
}

func (ms *MemoryStorage) Update(id int, task Task) {
	ms.tasks[id-1] = task
}

func (ms *MemoryStorage) Delete(id int) {
	ms.tasks = slices.Delete(ms.tasks, id-1, 1)
}

func main() {
	storage := &MemoryStorage{}

	storage.CreateTask(
		1,
		3,
		"Сделать лабораторную",
		"Написать трекер задач на Go",
		"Сдать до конца недели",
		25.5,
		time.Date(0, 1, 1, 14, 30, 0, 0, time.Local),
		time.Date(2026, time.October, 15, 0, 0, 0, 0, time.Local),
		false,
	)

	storage.CreateTask(
		2,
		4,
		"Сделать проект",
		"Ряды",
		"Сдать завтра",
		0,
		time.Date(0, 1, 1, 14, 30, 0, 0, time.Local),
		time.Date(2026, time.October, 15, 0, 0, 0, 0, time.Local),
		false,
	)
	fmt.Println(storage)
	task1 := storage.GetTask(1)
	fmt.Println(task1)

	storage.Delete(1)
	fmt.Println(storage)
	storage.Update(1, task1)
	fmt.Println(storage)
}
