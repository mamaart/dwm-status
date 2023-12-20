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

func Stream(chan<- error) (<-chan models.Time, error) {
	ch := make(chan models.Time)
	go stream(ch)
	return ch, nil
}

func stream(output chan<- models.Time) {
	state := Clock
	for {
		t := time.Now()
		switch state {
		case Clock:
			output <- models.Clock(t.Format("15:04"))
			state = Date
		case Date:
			output <- models.Calendar(t.Format("02 Jan 2006"))
			state = Weeknumber
		case Weeknumber:
			_, w := t.ISOWeek()
			output <- models.WeekNo(fmt.Sprintf("Week: %d", w))
			state = Day
		case Day:
			output <- models.Day(t.Weekday().String())
			state = Clock
		}
		time.Sleep(time.Second * 10)
	}
}
