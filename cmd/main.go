package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"project/internal/app"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	app := app.BuildApp()
	defer app.Close()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	// testing
	//////checking
	// var first, second int
	// err = pool.QueryRow(ctx, `SELECT 1, 2`).Scan(&first, &second)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(first, second)

	// store := storage.NewPostgresStorage(pool)

	// task := dto.Task{
	// 	Name:        "Сделать лабораторную",
	// 	Priority:    3,
	// 	Discription: "Написать трекер задач",
	// 	Time:        time.Date(0, 1, 1, 14, 30, 0, 0, time.UTC),
	// 	Date:        time.Date(2026, time.October, 15, 0, 0, 0, 0, time.UTC),
	// 	Status:      dto.StatusToDo,
	// }

	// if err := store.CreateTask(ctx, task); err != nil {
	// 	log.Fatal(err)
	// }

	// got, err := store.GetTask(ctx, 2)
	// fmt.Println(got, err)

	// task.Name = "Новое имя"
	// fmt.Println(store.UpdateTask(ctx, 2, task))

	// fmt.Println(store.DeleteTask(ctx, 2))
	// fmt.Println(store.DeleteTask(ctx, 2))

	if err := app.Start(":8080"); err != nil {
		panic(fmt.Sprintf("Ошибка при запуске приложения: %v", err))
	}

	// 1, 2
}
