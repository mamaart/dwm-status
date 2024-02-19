package volumemodule

import "fmt"

type Volume int

func (v Volume) String() string {
	return fmt.Sprintf(" 🎵 %d%% ", v)
}
