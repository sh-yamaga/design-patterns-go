package season

import "time"

type Direction int

type Season string

const (
	North Direction = iota
	South
)

const (
	Spring Season = "Spring"
	Summer Season = "Summmer"
	Fall   Season = "Fall"
	Winter Season = "Winter"
)

type MonthSeason struct {
	direction Direction
	month     time.Month
}

func NewMonthSeason(direction Direction, month time.Month) *MonthSeason {
	return &MonthSeason{
		direction: direction,
		month:     month,
	}
}

func (h MonthSeason) Month() string {
	return h.month.String()
}

func (h MonthSeason) Season() Season {
	switch h.direction {
	case North:
		switch h.month {
		case time.March, time.April, time.May:
			return Spring
		case time.June, time.July, time.August:
			return Summer
		case time.September, time.October, time.November:
			return Fall
		case time.December, time.January, time.February:
			return Winter
		}
	case South:
		switch h.month {
		case time.March, time.April, time.May:
			return Fall
		case time.June, time.July, time.August:
			return Winter
		case time.September, time.October, time.November:
			return Spring
		case time.December, time.January, time.February:
			return Summer
		}
	}
	return ""
}
