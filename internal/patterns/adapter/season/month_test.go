package season

import (
	"testing"
	"time"
)

func TestMonth(t *testing.T) {
	now := time.Now()
	m := NewMonth(now)

	if m == nil {
		t.Fatal("NewMonth returned nil")
	}

	got := m.Month
	want := now.Month()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
