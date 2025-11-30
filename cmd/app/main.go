package main

import (
	"awesomeProject/internal/db"
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/taskService"
	"awesomeProject/internal/userService"
	"awesomeProject/internal/web/tasks"
	"awesomeProject/internal/web/users"
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
	tskService := taskService.NewTasksService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(tskService)

	userRepo := userService.NewUserRepository(database)
	usrService := userService.NewUserService(userRepo, tskService)
	userHandlers := handlers.NewUserHandler(usrService)

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	strictHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, strictHandler)

	strictUserHandler := users.NewStrictHandler(userHandlers, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
