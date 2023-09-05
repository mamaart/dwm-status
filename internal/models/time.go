package models

type Time struct {
	Calendar, Clock string
}

type Calendar string

func (c Calendar) String() string {
	return "📅 " + string(c)
}

type Clock string

func (c Clock) String() string {
	return "🕒 " + string(c)
}
