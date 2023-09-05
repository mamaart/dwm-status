package application

import (
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

func (a *App) timeLoop() models.Time {
	for {
		currentTime := time.Now()
		a.time <- models.Time{
			Calendar: currentTime.Format("2006-01-02"),
			Clock:    currentTime.Format("15:04"),
		}
		time.Sleep(time.Minute)
	}
}
