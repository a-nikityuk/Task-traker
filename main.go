package main

import (
	"fmt"
	"slices"
	"time"
)

type Task struct {
	Id int
	Priority int
	Name string
	Discrip string
	ProgressComplete float64
	Time time.Time
	Date time.Time
	IsDo bool
}

type Storage interface {
	CreateTask(id, priority int, name, discrip string, pComplete float64, time, date time.Time, is_Do bool)
	GetTask(id int) Task
	Update(id int, task Task)
	Delete(id int)
}

type MemoryStorage struct {
	Tasks []Task
}

func (ms *MemoryStorage) CreateTask(id, priority int, name, discrip string, pComplete float64, time, date time.Time, is_Do bool) {
	task := Task{
		Id: id,
		Priority: priority,
		Name: name,
		Discrip: discrip,
		ProgressComplete: pComplete,
		Time: time,
		Date: date,
		IsDo: is_Do,
	}
	ms.Tasks = append(ms.Tasks, task)
}

func (ms *MemoryStorage) GetTask(id int) Task {
	getTask := ms.Tasks[id-1]
	return getTask
}

func (ms *MemoryStorage) Update(id int, task Task) {
	ms.Tasks[id-1] = task
}

func (ms *MemoryStorage) Delete(id int) {
	ms.Tasks = slices.Delete(ms.Tasks, id-1, 1)
}

func main() {
	storage := &MemoryStorage{}

	storage.CreateTask(
		1,
		3,
		"Сделать лабораторную",
		"Написать трекер задач на Go",
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
