package storage

import (
	"context"
	"fmt"
	"project/internal/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
	CreateTask(context.Context, dto.Task) error
	GetTask(context.Context, int) (dto.Task, error)
	UpdateTask(context.Context, int, dto.Task) error
	DeleteTask(context.Context, int) error
}

type PostgresStorage struct {
	db *pgxpool.Pool
}

func NewPostgresStorage(db *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{
		db: db,
	}
}

type MemoryStorage struct {
	tasksStorage map[int]dto.Task
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasksStorage: make(map[int]dto.Task),
	}
}

func (ps *PostgresStorage) CreateTask(ctx context.Context, task dto.Task) error {

	_, err := ps.db.Exec(ctx,
		`INSERT INTO tasks (name, priority, description, progress_complete, task_time, task_date, status)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		task.Name, task.Priority, task.Discription, task.ProgressComplete, task.Time, task.Date, task.Status,
	)

	return err
}

func (ps *PostgresStorage) GetTask(ctx context.Context, id int) (dto.Task, error) {

	var task dto.Task

	err := ps.db.QueryRow(
		ctx,
		`SELECT id, name, priority, description, progress_complete, task_time, task_date, status
		 FROM tasks
		 WHERE id = $1`, id,
	).Scan(&task.Id, &task.Name, &task.Priority, &task.Discription,
		&task.ProgressComplete, &task.Time, &task.Date, &task.Status)

	if err != nil {
		return dto.Task{}, err
	}

	return task, err
}

func (ps *PostgresStorage) UpdateTask(ctx context.Context, id int, task dto.Task) error {

	tag, err := ps.db.Exec(ctx,
		`UPDATE tasks
		SET name = $1, priority = $2, description = $3, progress_complete = $4,
		     task_time = $5, task_date = $6, status = $7
		WHERE id = $8`,
		task.Name, task.Priority, task.Discription, task.ProgressComplete,
		task.Time, task.Date, task.Status, id,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("task %d not found", id)
	}
	return nil
}

func (ps *PostgresStorage) DeleteTask(ctx context.Context, id int) error {
	tag, err := ps.db.Exec(ctx,
		`DELETE FROM tasks
		WHERE id = $1`, id,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("task %d not found", id)
	}
	return nil
}

func (ms *MemoryStorage) CreateTask(ctx context.Context, task dto.Task) error {
	if ms.tasksStorage == nil {
		return fmt.Errorf("problems with MemoryStorage")
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

func (ms *MemoryStorage) UpdateTask(ctx context.Context, id int, task dto.Task) error {
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

func (ms *MemoryStorage) DeleteTask(ctx context.Context, id int) error {
	if ms.tasksStorage == nil {
		return fmt.Errorf("Problems with MemoryStorage")
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	if _, ok := ms.tasksStorage[id]; !ok {
		return fmt.Errorf("task %d not found", id)
	}

	delete(ms.tasksStorage, id)
	return nil
}
