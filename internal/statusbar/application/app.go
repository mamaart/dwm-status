package application

import "github.com/mamaart/statusbar/internal/models"

type App struct {
	iface      chan models.IFace
	text       chan string
	time       chan models.Time
	volume     chan int
	battery    chan models.Battery
	brightness chan int
}

func New() *App {
	return &App{
		iface:      make(chan models.IFace),
		text:       make(chan string),
		time:       make(chan models.Time),
		volume:     make(chan int),
		battery:    make(chan models.Battery),
		brightness: make(chan int),
	}
}

func (a *App) Run(ch chan<- models.State) {
	go a.textLoop()
	go a.ifLoop()
	go a.timeLoop()
	go a.volumeLoop()
	go a.batteryLoop()
	go a.brightnessLoop()

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
		case b := <-a.brightness:
			s.Brightness = b
		}
		ch <- s
	}
}
