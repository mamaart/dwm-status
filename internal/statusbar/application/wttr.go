package application

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

func (a *App) wttrLoop() {
	for {
		w, e := GetWttr()
		if e != nil {
			fmt.Println(e)
			time.Sleep(time.Second * 10)
		}
		a.wttr <- w
		time.Sleep(time.Second * 2)

	}
}

func GetWttr() (b models.Wttr, err error) {
	r, e := http.Get(`https://wttr.in/Copenhagen?format="+%t"`)
	if e != nil {
		return b, e
	}
	bts, e := io.ReadAll(r.Body)
	if e != nil {
		return b, e
	}
	return models.Wttr{Temp: string(bts)}, nil
}
