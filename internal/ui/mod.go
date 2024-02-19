package ui

import (
	"strings"

	"github.com/mamaart/statusbar/internal/ports"
	"github.com/mamaart/statusbar/pkg/bar"
)

func Run(
	readers ...ports.UniStreamer,
) {
	state := make(chan []byte)

	go func() {
		chunks := make([]string, len(readers))
		for {
			for i, ch := range readers {
				select {
				case data := <-ch.Reader():
					chunks[i] = string(data)
					state <- []byte(strings.Join(chunks, "|"))
				default:
				}
			}
		}
	}()

	bar := bar.New()
	for x := range state {
		bar.Update(x)
	}
}
