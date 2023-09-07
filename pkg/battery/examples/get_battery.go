package main

import (
	"fmt"
	"log"

	"github.com/mamaart/statusbar/pkg/battery"
)

func main() {
	battery, err := battery.Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Battery%+v", battery)
}
