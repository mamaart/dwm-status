package models

import (
	"fmt"
	"strings"
)

type State struct {
	Iface   IFace
	Text    string
	Time    Time
	Volume  float32
	Battery Battery
}

func (s State) Bytes() []byte {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("📡 %s | ", s.Iface.Addr))
	sb.WriteString(fmt.Sprintf("🎵 %d | ", int(s.Volume*100)))
	if s.Battery.Charging {
		sb.WriteString(fmt.Sprintf("💡 %d | ", s.Battery.Capacity))
	} else if s.Battery.Capacity < 50 {
		sb.WriteString(fmt.Sprintf("🪫 %d | ", s.Battery.Capacity))
	} else {
		sb.WriteString(fmt.Sprintf("🔋 %d | ", s.Battery.Capacity))
	}
	sb.WriteString(fmt.Sprintf("📅 %s | ", s.Time.Calendar))
	sb.WriteString(fmt.Sprintf("🕒 %s | ", s.Time.Clock))
	sb.WriteString(fmt.Sprintf("▶%s◀ ", s.Text))
	return []byte(sb.String())
}
