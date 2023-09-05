package application

import (
	"strings"
	"time"

	"github.com/mamaart/statusbar/internal/models"
	"github.com/mamaart/statusbar/internal/statusbar/database"
	"github.com/mamaart/statusbar/internal/statusbar/server"
)

const windowWidth = 40

func (a *App) textLoop() {
	var (
		ch     = make(chan []models.Task)
		db     = database.New()
		server = server.New(db)
	)

	go db.Run(ch)
	go server.Run()

	tasks := db.List()
	for {
		if len(tasks) != 0 {
			tasks = a.itterate(ch, tasks)
		} else {
			a.writeText(0, strings.Repeat(" ", 40))
			tasks = <-ch
		}
	}
}

func (a *App) itterate(ch <-chan []models.Task, tasks []models.Task) []models.Task {
	text := tasksToText(tasks)
	for i := 0; i <= len(text); i++ {
		if t, ok := check(ch); ok {
			return t
		}
		a.writeText(i, text)
	}
	return tasks
}

func (a *App) writeText(i int, text string) {
	start := i % len(text)
	end := start + windowWidth

	if end > len(text) {
		end = len(text)
	}

	if end-start < windowWidth {
		remaining := windowWidth - (end - start)
		a.text <- text[start:] + text[:remaining]
	} else {
		a.text <- text[start:end]
	}
	time.Sleep(time.Millisecond * 200)
}

func check(ch <-chan []models.Task) ([]models.Task, bool) {
	select {
	case t := <-ch:
		return t, true

	default:
		return nil, false
	}
}

func tasksToText(tasks []models.Task) string {
	var sb strings.Builder

	if len(tasks) == 0 {
		return ""
	}

	for _, t := range tasks {
		sb.WriteString(t.String())
		sb.WriteString(" ")
	}

	text := sb.String()

	if len(text) < 40 {
		pad := strings.Repeat(" ", 40-len(text))
		text = pad + text
	}

	return text
}
