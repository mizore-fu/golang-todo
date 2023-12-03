package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}

func (t *Task) Validate() error {
    return validation.ValidateStruct(t,
		validation.Field(
			&t.Name,
			validation.Required.Error("name is required"),
		),
	)
}
