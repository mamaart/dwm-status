package models

import (
	"fmt"
)

type State struct {
	Iface      IFace
	Text       Text
	Clock      Time
	Volume     Volume
	Battery    Battery
	Brightness Brightness
	Wttr       Wttr
	Disk       Disk
}

func (s State) Bytes() []byte {
	return []byte(fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s | %s ",
		s.Clock,
		s.Iface,
		s.Brightness,
		s.Volume,
		s.Battery,
		&s.Wttr,
		s.Disk,
		s.Text,
	))
}
