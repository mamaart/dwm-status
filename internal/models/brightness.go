package models

import "fmt"

type Brightness int

func (b Brightness) String() string {
	return fmt.Sprintf("🔆 %d%%", b)
}
