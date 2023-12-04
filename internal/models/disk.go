package models

type Disk string

func (d Disk) String() string {
	return "💾 " + string(d)
}
