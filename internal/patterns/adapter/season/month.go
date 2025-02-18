package season

import "time"

// Month is a struct that wraps a time.Month
type Month struct {
	Month time.Month
}

// NewMonth creates a new Month with the given time
func NewMonth(time time.Time) *Month {
	return &Month{Month: time.Month()}
}
