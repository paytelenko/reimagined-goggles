package main

import (
	"awesomeProject/internal/db"
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/taskService"
	"awesomeProject/internal/web/tasks"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitBD()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	e := echo.New()

	taskRepo := taskService.NewTaskRepository(database)
	service := taskService.NewTasksService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(service)

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	strictHandler := tasks.NewStrictHandler(taskHandlers, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
