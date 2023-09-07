package main

import (
	"fmt"
	"log"

	"github.com/mamaart/statusbar/pkg/brightness"
)

func main() {
	brightness, err := brightness.Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Brightness(%+v%%)\n", brightness)
}
