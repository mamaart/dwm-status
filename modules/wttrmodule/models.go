package wttrmodule

type Wttr struct {
	Temp string
}

func (wttr *Wttr) String() string {
	return wttr.Temp
}
