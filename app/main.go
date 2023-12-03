package main

import (
	"app/model"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Tasks struct {
	tasks []*model.Task
}

func (ts *Tasks) GetAllTasks() []*model.Task {
	return ts.tasks
}

func (ts *Tasks) AddTask(c echo.Context) error {
	addedTask := new(model.Task)
	if err := c.Bind(addedTask); err != nil {
		return err
	}
	addedTask.ID = "taskid-" + uuid.NewString()
	addedTask.Completed = false

	if err := addedTask.Validate(); err != nil {
		return err
	}

	ts.tasks = append(ts.tasks, addedTask)
	return nil
}

func (ts *Tasks) DeleteTask(id string) error {
	position := -1
	for i, task := range ts.tasks {
		if id == task.ID {
			position = i
			break
		}
	}
	if position == -1 {
		return errors.New("削除対象のタスクが見つかりませんでした。")
	}

	ts.tasks[position] = ts.tasks[len(ts.tasks) - 1]
	ts.tasks = ts.tasks[:len(ts.tasks) - 1]
	return nil
}

var tasks *Tasks = &Tasks{
	tasks: []*model.Task{
		{ID: "1", Name: "eat", Completed: false},
		{ID: "2", Name: "sleep", Completed: false},
	},
}



func main() {
	e := echo.New()
	e.GET("/tasks", GetAllTasksHandler)
	e.POST("/tasks", AddTaskHandler)
	e.DELETE("/tasks/:id", DeleteTaskHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

//GET /tasks
func GetAllTasksHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks.GetAllTasks())
}

//POST /tasks
//body: {name: "test"}
func AddTaskHandler(c echo.Context) error {
	if err := tasks.AddTask(c); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, nil)
}

//DELETE /tasks/taskid-12345
func DeleteTaskHandler(c echo.Context) error {
	id := c.Param("id")
	if err := tasks.DeleteTask(id); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
