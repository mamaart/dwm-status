package models

type Text string

func (t Text) String() string {
	return "▶" + string(t) + "◀ "
}
