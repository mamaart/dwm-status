package main

import (
	"fmt"
	"time"

	"github.com/mamaart/statusbar/internal/models"
	"github.com/mamaart/statusbar/internal/ports"
	"github.com/mamaart/statusbar/internal/statusbar/database"
	"github.com/mamaart/statusbar/pkg/tasks"
)

func main() {
	db := database.New()
	go feed(db)

	tm, err := tasks.New(tasks.Options{
		WindowWidth: 40,
		Database:    db,
		Delay:       time.Millisecond * 100,
	}).Stream(nil)
	if err != nil {
		fmt.Println(tm)
	}

	for e := range tm {
		fmt.Printf("\r%s", e)
	}
}

func feed(db ports.Database) {
	db.Add(models.Task{
		Description: "This is a task",
	})
}
