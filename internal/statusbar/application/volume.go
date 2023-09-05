package application

import (
	"fmt"

	"github.com/mafik/pulseaudio"
)

func (a *App) volumeLoop() {
	cli, err := pulseaudio.NewClient()
	if err != nil {
		panic(err)
	}

	v, _ := cli.Volume()
	a.volume <- int(v * 100)

	updates, err := cli.Updates()
	if err != nil {
		panic(err)
	}

	for range updates {
		v, err := cli.Volume()

		if err != nil {
			fmt.Println(err)
			continue
		}

		a.volume <- int(v * 100)
	}
}
