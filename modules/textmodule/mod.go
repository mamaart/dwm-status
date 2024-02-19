package textmodule

import (
	"errors"
	"time"
)

type Options struct {
	WindowWidth int
	Delay       time.Duration
}

type TextModule struct {
	windowWidth int
	delay       time.Duration
	input       chan<- byte
	output      <-chan []byte
	toggle      chan<- struct{}
}

func New(options Options) *TextModule {
	if options.Delay == time.Duration(0) {
		options.Delay = time.Millisecond * 500
	}

	output := make(chan []byte)
	input := make(chan byte)
	toggle := make(chan struct{})

	go func() {
		i := options.WindowWidth - 1

		array := make([]byte, options.WindowWidth)
		for i := range array {
			array[i] = 32
		}

		show := false

		for {
			for j := 0; j < options.WindowWidth-1; j++ {
				array[j] = array[j+1]
			}

			select {
			case <-toggle:
				show = !show

			case x := <-input:
				array[options.WindowWidth-1] = x
				time.Sleep(options.Delay)

			case <-time.After(options.Delay):
				array[options.WindowWidth-1] = 32
			}

			if show {
				output <- []byte(Text(array[:]).String())
			} else {
				output <- []byte{}
			}
			i = (i - 1) % options.WindowWidth
		}
	}()

	return &TextModule{
		input:  input,
		output: output,
		toggle: toggle,
	}
}

func (t *TextModule) Reader() <-chan []byte {
	return t.output
}

func (t *TextModule) Write(in []byte) (n int, err error) {
	for i, e := range in {
		select {
		case t.input <- e:
		case <-time.After(time.Second * 5):
			return i, errors.New("timeout when writing stream")
		}
	}
	return -1, nil
}

func (t *TextModule) Toggle() error {
	select {
	case t.toggle <- struct{}{}:
		return nil
	case <-time.After(time.Second * 5):
		return errors.New("timeout while toggeling text")
	}
}
