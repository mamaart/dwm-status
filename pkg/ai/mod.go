package ai

import (
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

type Options struct {
	WindowWidth int
	Delay       time.Duration
}

type Ai struct {
	windowWidth int
	delay       time.Duration
	ch          <-chan byte
}

func New(ch <-chan byte, options Options) *Ai {
	if options.Delay == time.Duration(0) {
		options.Delay = time.Millisecond * 500
	}
	return &Ai{
		windowWidth: options.WindowWidth,
		delay:       options.Delay,
		ch:          ch,
	}
}

func (ai *Ai) Stream(errCh chan<- error) (chan models.Text, error) {
	ch := make(chan models.Text)
	go ai.stream(ch)
	return ch, nil
}

func (ai *Ai) stream(output chan<- models.Text) {
	i := ai.windowWidth - 1

	array := make([]byte, ai.windowWidth)
	for i := range array {
		array[i] = 32
	}

	for {

		for j := 0; j < ai.windowWidth-1; j++ {
			array[j] = array[j+1]
		}

		select {
		case x := <-ai.ch:
			array[ai.windowWidth-1] = x
			time.Sleep(ai.delay)

		case <-time.After(ai.delay):
			array[ai.windowWidth-1] = 32
		}

		output <- models.Text(array[:])
		i = (i - 1) % ai.windowWidth
	}
}
