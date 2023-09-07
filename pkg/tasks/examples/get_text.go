package main

import (
	"fmt"

	"github.com/mamaart/statusbar/internal/statusbar/database"
	"github.com/mamaart/statusbar/internal/statusbar/server"
	"github.com/mamaart/statusbar/pkg/tasks"
)

func main() {
	db := database.New()

	go server.New(db).Run()

	tm, err := tasks.New(tasks.Options{
		WindowWidth: 40,
		Database:    db,
	}).Stream(nil)
	if err != nil {
		fmt.Println(tm)
	}

	for e := range tm {
		fmt.Println(e)
	}
}
