package model

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}
