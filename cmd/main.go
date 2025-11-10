package main

import (
	"awesomeProject/internal/db"
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/taskService"
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
	taskServise := taskService.NewTasksService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskServise)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/tasks", taskHandlers.PostTasks)
	e.GET("/tasks", taskHandlers.GetTasks)
	e.PATCH("/tasks/:id", taskHandlers.PatchTask)
	e.DELETE("/tasks/:id", taskHandlers.DeleteTask)

	e.Start("localhost:8080")
}
