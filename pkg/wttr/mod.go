package wttr

import (
	"io"
	"net/http"
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

func Get() (b models.Wttr, err error) {
	r, e := http.Get(`https://wttr.in/Copenhagen?format=%t`)
	if e != nil {
		return b, e
	}
	bts, e := io.ReadAll(r.Body)
	if e != nil {
		return b, e
	}
	return models.Wttr{Temp: string(bts)}, nil
}

func Stream(errch chan<- error) (<-chan models.Wttr, error) {
	ch := make(chan models.Wttr)
	go stream(ch, errch)
	return ch, nil
}

func stream(output chan<- models.Wttr, err chan<- error) {
	for {
		w, e := Get()
		if e != nil {
			if err != nil {
				err <- e
			}
			time.Sleep(time.Minute)
		} else {
			output <- w
			time.Sleep(time.Minute * 10)
		}
	}
}
