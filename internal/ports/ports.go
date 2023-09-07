package ports

import "github.com/mamaart/statusbar/internal/models"

type Database interface {
	Add(models.Task)
	Delete(int)
	List() []models.Task
	Stream() <-chan []models.Task
}
