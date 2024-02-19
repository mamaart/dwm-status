package batterymodule

import "fmt"

type Battery struct {
	Charging bool
	Capacity int
}

func (b Battery) String() string {
	if b.Charging {
		return fmt.Sprintf(" ⚡ %d ", b.Capacity)
	} else if b.Capacity < 50 {
		return fmt.Sprintf(" 🪫 %d ", b.Capacity)
	} else {
		return fmt.Sprintf(" 🔋 %d ", b.Capacity)
	}
}
