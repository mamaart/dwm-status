package tasks

import (
	"strings"
	"time"

	"github.com/mamaart/statusbar/internal/models"
	"github.com/mamaart/statusbar/internal/ports"
)

type Manager struct {
	windowWidth int
	database    ports.Database
}

type Options struct {
	WindowWidth int
	Database    ports.Database
}

func New(options Options) *Manager {
	return &Manager{
		windowWidth: options.WindowWidth,
		database:    options.Database,
	}
}

func (t *Manager) Stream(errch chan<- error) (<-chan models.Text, error) {
	ch := make(chan models.Text)
	go t.stream(ch)
	return ch, nil
}

func (t *Manager) stream(output chan<- models.Text) {
	tasks := t.database.List()
	input := t.database.Stream()

	for {
		if len(tasks) != 0 {
			tasks = t.itterate(input, output, tasks)
		} else {
			t.writeText(output, 0, strings.Repeat(" ", 40))
			tasks = <-input
		}
	}
}

func (a *Manager) itterate(
	input <-chan []models.Task,
	output chan<- models.Text,
	tasks []models.Task,
) []models.Task {
	text := tasksToText(tasks)
	for i := 0; i <= len(text); i++ {
		if t, ok := check(input); ok {
			return t
		}
		a.writeText(output, i, text)
	}
	return tasks
}

func (t *Manager) writeText(output chan<- models.Text, i int, text string) {
	start := i % len(text)
	end := start + t.windowWidth

	if end > len(text) {
		end = len(text)
	}

	if end-start < t.windowWidth {
		remaining := t.windowWidth - (end - start)
		output <- models.Text(text[start:] + text[:remaining])
	} else {
		output <- models.Text(text[start:end])
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
