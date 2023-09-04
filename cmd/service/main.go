package main

import (
	"github.com/mamaart/statusbar/internal/application"
	"github.com/mamaart/statusbar/internal/bar"
	"github.com/mamaart/statusbar/internal/models"
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
