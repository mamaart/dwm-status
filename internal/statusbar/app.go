package statusbar

import (
	"github.com/mamaart/statusbar/internal/models"
	"github.com/mamaart/statusbar/pkg/battery"
	"github.com/mamaart/statusbar/pkg/brightness"
	"github.com/mamaart/statusbar/pkg/datetime"
	"github.com/mamaart/statusbar/pkg/network"
	"github.com/mamaart/statusbar/pkg/tasks"
	"github.com/mamaart/statusbar/pkg/volume"
	"github.com/mamaart/statusbar/pkg/wttr"
)

type StatusBar struct {
	iface      <-chan models.IFace
	text       <-chan models.Text
	time       <-chan models.Time
	volume     <-chan models.Volume
	battery    <-chan models.Battery
	brightness <-chan models.Brightness
	wttr       <-chan models.Wttr
}

func New() (*StatusBar, error) {
	iface, err := network.Stream(nil)
	if err != nil {
		return nil, err
	}
	time, err := datetime.Stream(nil)
	if err != nil {
		return nil, err
	}
	volume, err := volume.Stream(nil)
	if err != nil {
		return nil, err
	}
	battery, err := battery.Stream(nil)
	if err != nil {
		return nil, err
	}
	brightness, err := brightness.Stream(nil)
	if err != nil {
		return nil, err
	}
	wttr, err := wttr.Stream(nil)
	if err != nil {
		return nil, err
	}

	text, err := tasks.New(40).Stream(nil)
	if err != nil {
		return nil, err
	}

	return &StatusBar{
		text:       text,
		iface:      iface,
		time:       time,
		volume:     volume,
		battery:    battery,
		brightness: brightness,
		wttr:       wttr,
	}, nil
}

func (a *StatusBar) Run(ch chan<- models.State) {
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
		case w := <-a.wttr:
			s.Wttr = w
		}
		ch <- s
	}
}
