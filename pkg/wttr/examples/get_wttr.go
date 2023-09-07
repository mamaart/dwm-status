package main

import (
	"fmt"
	"log"

	"github.com/mamaart/statusbar/pkg/wttr"
)

func main() {
	wttr, err := wttr.Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wttr: %+v", wttr)
}
