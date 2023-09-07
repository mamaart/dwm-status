package main

import (
	"fmt"

	"github.com/mamaart/statusbar/pkg/datetime"
)

func main() {
	timedate := datetime.Get()
	fmt.Printf("Time: %+v\n", timedate.Clock.String())
	fmt.Printf("Date: %+v\n", timedate.Calendar.String())
}
