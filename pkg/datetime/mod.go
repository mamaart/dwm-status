package datetime

import (
	"fmt"
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

type WatchFace uint

const (
	Clock WatchFace = iota
	Date
	Weeknumber
	Day
)

func Stream(control <-chan struct{}, _ chan<- error) (<-chan models.Time, error) {
	ch := make(chan models.Time)
	go stream(control, ch)
	return ch, nil
}

func stream(control <-chan struct{}, output chan<- models.Time) {
	state := Clock
	output <- models.Clock(time.Now().Format("15:04"))
	for {
		select {
		case <-control:
			t := time.Now()
			switch state {
			case Clock:
				state = Date
				output <- models.Calendar(t.Format("02 Jan 2006"))
			case Date:
				state = Weeknumber
				_, w := t.ISOWeek()
				output <- models.WeekNo(fmt.Sprintf("Week: %d", w))
			case Weeknumber:
				state = Day
				output <- models.Day(t.Weekday().String())
			case Day:
				state = Clock
				output <- models.Clock(t.Format("15:04"))
			}
		case <-time.After(time.Second * 5):
			t := time.Now()
			switch state {
			case Clock:
				output <- models.Clock(t.Format("15:04"))
			case Date:
				output <- models.Calendar(t.Format("02 Jan 2006"))
			case Weeknumber:
				_, w := t.ISOWeek()
				output <- models.WeekNo(fmt.Sprintf("Week: %d", w))
			case Day:
				output <- models.Day(t.Weekday().String())
			}
		}
	}
}
