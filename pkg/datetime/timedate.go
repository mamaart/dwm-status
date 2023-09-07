package datetime

import (
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

func Get() models.Time {
	currentTime := time.Now()
	return models.Time{
		Calendar: models.Calendar(currentTime.Format("01/02 2006")),
		Clock:    models.Clock(currentTime.Format("15:04")),
	}
}

func Stream(chan<- error) (chan models.Time, error) {
	ch := make(chan models.Time)
	go stream(ch)
	return ch, nil
}

func stream(output chan<- models.Time) {
	for {
		output <- Get()
		time.Sleep(time.Minute)
	}
}
