package wttrmodule

import (
	"io"
	"log"
	"net/http"
	"time"
)

type WttrModule struct {
	output <-chan []byte
}

func New() *WttrModule {
	output := make(chan []byte)
	go func() {
		for {
			w, err := Get()
			if err != nil {
				log.Println(err)
				time.Sleep(time.Minute)
			} else {
				output <- []byte(w.String())
				time.Sleep(time.Minute * 10)
			}
		}
	}()
	return &WttrModule{output: output}
}

func (w *WttrModule) Reader() <-chan []byte {
	return w.output
}

func Get() (b Wttr, err error) {
	r, e := http.Get(`https://wttr.in/Copenhagen?format=%t`)
	if e != nil {
		return b, e
	}
	bts, e := io.ReadAll(r.Body)
	if e != nil {
		return b, e
	}
	return Wttr{Temp: string(bts)}, nil
}
