package ports

import "github.com/mamaart/statusbar/internal/models"

type Text interface {
	Stream(chan<- error) (<-chan models.Text, error)
}
