package main

import (
	"github.com/mamaart/statusbar/internal/models"
	"github.com/mamaart/statusbar/internal/statusbar/application"
	"github.com/mamaart/statusbar/pkg/bar"
)

func main() {
	bar := bar.New()
	ch := make(chan models.State)
	app := application.New()
	go app.Run(ch)
	for x := range ch {
		bar.Update(x.Bytes())
	}
}
