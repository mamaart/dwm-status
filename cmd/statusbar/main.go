package main

import (
	"log"

	"github.com/mamaart/statusbar/internal/models"
	"github.com/mamaart/statusbar/internal/statusbar"
	"github.com/mamaart/statusbar/internal/statusbar/server"
	"github.com/mamaart/statusbar/pkg/bar"
)

func main() {
	bar := bar.New()
	output := make(chan models.State)
	input := make(chan byte)

	app, err := statusbar.New(input)
	if err != nil {
		log.Fatal(err)
	}

	go server.Run(input)
	go app.Run(output)

	for x := range output {
		bar.Update(x.Bytes())
	}
}
