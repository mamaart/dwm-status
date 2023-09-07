package main

import (
	"log"

	"github.com/mamaart/statusbar/internal/models"
	"github.com/mamaart/statusbar/internal/statusbar"
	"github.com/mamaart/statusbar/pkg/bar"
)

func main() {
	bar := bar.New()
	ch := make(chan models.State)
	app, err := statusbar.New()
	if err != nil {
		log.Fatal(err)
	}
	go app.Run(ch)
	for x := range ch {
		bar.Update(x.Bytes())
	}
}
