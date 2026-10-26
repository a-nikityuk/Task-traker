package main

import (
	"fmt"
	"project/internal/app"
)

func main() {
	app := app.BuildApp()
	defer app.Close()

	if err := app.Start(":8080"); err != nil {
		panic(fmt.Sprintf("Ошибка при запуске приложения: %v", err))
	}
}
