package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}

func (task *Task) Validate() error {
    return validation.ValidateStruct(task,
		validation.Field(
			&task.Name,
			validation.Required.Error("name is required"),
		),
	)
}
