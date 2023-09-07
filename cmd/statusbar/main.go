package main

import (
	"log"

	"github.com/mamaart/statusbar/internal/models"
	"github.com/mamaart/statusbar/internal/statusbar"
	"github.com/mamaart/statusbar/internal/statusbar/database"
	"github.com/mamaart/statusbar/internal/statusbar/server"
	"github.com/mamaart/statusbar/pkg/bar"
)

func main() {
	bar := bar.New()
	ch := make(chan models.State)
	db := database.New()

	app, err := statusbar.New(db)
	if err != nil {
		log.Fatal(err)
	}

	go server.New(db).Run()
	go app.Run(ch)

	for x := range ch {
		bar.Update(x.Bytes())
	}
}
