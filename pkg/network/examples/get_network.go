package main

import (
	"fmt"
	"log"

	"github.com/mamaart/statusbar/pkg/network"
)

func main() {
	net, err := network.Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Network: %+v\n", net)
}
