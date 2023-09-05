package models

import (
	"fmt"
	"strings"
)

type State struct {
	Iface      IFace
	Text       string
	Time       Time
	Volume     int
	Battery    Battery
	Brightness int
}

func (s State) Bytes() []byte {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("📡 %s%% | ", s.Iface.Addr))
	sb.WriteString(fmt.Sprintf("☀️ %d%% | ", s.Brightness))
	sb.WriteString(fmt.Sprintf("🎵 %d%% | ", s.Volume))
	if s.Battery.Charging {
		sb.WriteString(fmt.Sprintf("⚡ %d | ", s.Battery.Capacity))
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
