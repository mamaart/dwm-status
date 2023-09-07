package main

import (
	"fmt"

	"github.com/mamaart/statusbar/pkg/tasks"
)

func main() {
	tm, err := tasks.New(40).Stream(nil)
	if err != nil {
		fmt.Println(tm)
	}

	for e := range tm {
		fmt.Println(e)
	}
}
