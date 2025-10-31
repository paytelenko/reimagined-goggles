package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func postTask(c echo.Context) error {
	var requestBody requestBody
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	task = requestBody.Task
	return c.JSON(http.StatusAccepted, task)
}

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello, "+task)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/task", postTask)
	e.GET("/task", getTask)

	e.Start("localhost:8080")
}
