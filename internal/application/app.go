package application

import "github.com/mamaart/statusbar/internal/models"

type App struct {
	iface   chan models.IFace
	text    chan string
	time    chan models.Time
	volume  chan float32
	battery chan models.Battery
}

func New() *App {
	return &App{
		iface:   make(chan models.IFace),
		text:    make(chan string),
		time:    make(chan models.Time),
		volume:  make(chan float32),
		battery: make(chan models.Battery),
	}
}

func (a *App) Run(ch chan<- models.State) {
	go a.textLoop()
	go a.ifLoop()
	go a.timeLoop()
	go a.volumeLoop()
	go a.batteryLoop()

	var s models.State
	for {
		select {
		case iface := <-a.iface:
			s.Iface = iface
		case txt := <-a.text:
			s.Text = txt
		case t := <-a.time:
			s.Time = t
		case v := <-a.volume:
			s.Volume = v
		case b := <-a.battery:
			s.Battery = b
		}
		ch <- s
	}
}
