package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}

var tasks []*Task = []*Task{}

func main() {
	tasks = append(tasks, &Task{ID: "1", Name: "eat", Completed: false})
	tasks = append(tasks, &Task{ID: "2", Name: "sleep", Completed: false})

	e := echo.New()
	e.GET("/tasks", getAllTasks)
	e.Logger.Fatal(e.Start(":8080"))
}

func getAllTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}
