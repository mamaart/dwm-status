package models

import (
	"fmt"
)

type State struct {
	Iface      IFace
	Text       Text
	Time       Time
	Volume     Volume
	Battery    Battery
	Brightness Brightness
	Wttr       Wttr
}

func (s State) Bytes() []byte {
	return []byte(fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s | %s",
		s.Iface,
		s.Brightness,
		s.Volume,
		s.Battery,
		s.Time.Calendar,
		s.Time.Clock,
		&s.Wttr,
		s.Text,
	))
}
