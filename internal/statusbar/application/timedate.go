package application

import (
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

func (a *App) timeLoop() models.Time {
	for {
		currentTime := time.Now()
		a.time <- models.Time{
			Calendar: models.Calendar(currentTime.Format("01/02 2006")),
			Clock:    models.Clock(currentTime.Format("15:04")),
		}
		time.Sleep(time.Minute)
	}
}
