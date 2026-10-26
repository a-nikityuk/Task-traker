package storage

import (
	"context"
	"fmt"
	"project/internal/dto"
)

type Storage interface {
	CreateTask(context.Context, dto.Task) error
	GetTask(context.Context, int) (dto.Task, error)
	Update(context.Context, int, dto.Task) error
	Delete(context.Context, int) error
}

type MemoryStorage struct {
	tasksStorage map[int]dto.Task
}

func (ms *MemoryStorage) CreateTask(ctx context.Context, task dto.Task) error {
	if ms.tasksStorage == nil {
		return fmt.Errorf("Problems with MemoryStorage")
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	ms.tasksStorage[task.Id] = task
	return nil
}

func (ms *MemoryStorage) GetTask(ctx context.Context, id int) (dto.Task, error) {
	if ms.tasksStorage == nil {
		return dto.Task{}, fmt.Errorf("Problems with MemoryStorage")
	}
	if err := ctx.Err(); err != nil {
		return dto.Task{}, err
	}
	return ms.tasksStorage[id], nil
}

func (ms *MemoryStorage) Update(ctx context.Context, id int, task dto.Task) error {
	if ms.tasksStorage == nil {
		return fmt.Errorf("Problems with MemoryStorage")
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	if _, ok := ms.tasksStorage[id]; !ok {
		return fmt.Errorf("task %d not found", id)
	}

	ms.tasksStorage[id] = task
	return nil
}

func (ms *MemoryStorage) Delete(ctx context.Context, id int) error {
	if ms.tasksStorage == nil {
		return fmt.Errorf("Problems with MemoryStorage")
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	if _, ok := ms.tasksStorage[id]; !ok {
		return fmt.Errorf("task %d not found", id)
	}

	ms.tasksStorage[id] = dto.Task{}
	return nil
}
