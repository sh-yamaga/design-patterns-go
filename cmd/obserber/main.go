package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/observer"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/observer/display"
)

func main() {
	weather := &observer.Weather{}

	// create display(Observer)
	currentDisplay := &display.CurrentDisplay{}

	// add Observers
	weather.AddObserver(currentDisplay)

	fmt.Println("=== Initial Measuremant ===")
	weather.SetMeasurements(25.0, 65.5, 1013.1)
	// Output:
	// === Initial Measuremant ===
	// ▼ Current Weather
	//   Temp :        25.0℃
	//   Humid:        65.5%
	//   Press:        1013.1hPa
	// ---------------------------
}
