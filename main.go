package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID        string `json:"id"`
	Objective string `json:"objective"`
	Status    string `json:"status"`
}

var tasks = []Task{}

type requestBody struct {
	Objective string `json:"objective"`
}

func postTask(c echo.Context) error {
	var requestBody requestBody
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	task := Task{
		ID:        uuid.NewString(),
		Objective: requestBody.Objective,
		Status:    "in progress",
	}
	tasks = append(tasks, task)
	return c.JSON(http.StatusCreated, task)
}

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")
	var requestBody requestBody
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = "done"
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task not found"})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Calculation not found"})
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/tasks", postTask)
	e.GET("/tasks", getTasks)
	e.PATCH("/tasks", patchTask)
	e.DELETE("/tasks", deleteTask)

	e.Start("localhost:8080")
}
