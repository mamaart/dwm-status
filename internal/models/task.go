package models

import "fmt"

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func (t *Task) String() string {
	return fmt.Sprintf("[%d] %s", t.Id, t.Description)
}
