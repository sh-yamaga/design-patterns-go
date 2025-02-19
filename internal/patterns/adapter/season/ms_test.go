package season

import (
	"testing"
	"time"
)

func TestSeasonSummer(t *testing.T) {
	au := time.Date(2025, time.August, 1, 0, 0, 0, 0, time.UTC)

	nh := NewMonthSeason(North, NewMonth(au).Month)

	got := nh.Season()
	want := Summer

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
