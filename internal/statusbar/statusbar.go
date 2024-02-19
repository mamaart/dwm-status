package statusbar

import "github.com/mamaart/statusbar/internal/ports"

type Module struct {
	input <-chan []byte
	value []byte
}

type StatusBar struct {
	modules []Module
}

func New(streamers ...ports.UniStreamer) *StatusBar {
	var modules []Module
	for _, e := range streamers {
		modules = append(modules, Module{input: e.Reader()})
	}
	return &StatusBar{modules}
}

func (a *StatusBar) AddModule(s ports.UniStreamer) {
	a.modules = append(a.modules, Module{input: s.Reader(), value: []byte{}})
}

func (a *StatusBar) StartAsync() <-chan []byte {
	output := make(chan []byte)
	go func() {
		for i, ch := range a.modules {
			select {
			case data := <-ch.input:
				a.modules[i].value = data
			default:
			}
		}
	}()
	return output
}
