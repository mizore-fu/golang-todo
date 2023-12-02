package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}

func NewTask() *Task {
	return &Task{}
}

type Tasks struct {
	tasks []*Task
}

func (ts *Tasks) GetAllTasks() []*Task {
	return ts.tasks
}

func (ts *Tasks) AddTask(c echo.Context) error {
	addedTask := NewTask()
	//TODO: Bindの使用について要検討({"bad": "testing"} このようなbodyの対策)
	if err := c.Bind(addedTask); err != nil {
		return err
	}

	id := "taskid-" + uuid.NewString()
	addedTask.ID = id
	ts.tasks = append(ts.tasks, addedTask)
	return nil
}

var tasks *Tasks = &Tasks{
	tasks: []*Task{
		{ID: "1", Name: "eat", Completed: false},
		{ID: "2", Name: "sleep", Completed: false},
	},
}



func main() {
	e := echo.New()
	e.GET("/tasks", GetAllTasksHandler)
	e.POST("/tasks", AddTaskHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

//GET /tasks
func GetAllTasksHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks.GetAllTasks())
}

//POST /tasks
//body: {name: "test"}
func AddTaskHandler(c echo.Context) error {
	err := tasks.AddTask(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, nil)
}
