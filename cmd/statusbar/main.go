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
	clockstate := make(chan struct{})

	app, err := statusbar.New(input, clockstate)
	if err != nil {
		log.Fatal(err)
	}

	go server.NewServer(input, clockstate).ListenAndServe()
	go app.Run(output)

	for x := range output {
		bar.Update(x.Bytes())
	}
}
