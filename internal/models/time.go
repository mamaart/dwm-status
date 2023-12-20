package models

type Time interface {
	String() string
}

type Calendar string

func (c Calendar) String() string {
	return "📅 " + string(c)
}

type Clock string

func (c Clock) String() string {
	return "🕒 " + string(c)
}

type WeekNo string

func (w WeekNo) String() string {
	return "📅 " + string(w)
}

type Day string

func (d Day) String() string {
	return "📅 " + string(d)
}
