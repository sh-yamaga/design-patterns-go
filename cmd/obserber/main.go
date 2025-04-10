package main

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/observer"
	"github.com/yamaga-shu/design-patterns-go/internal/patterns/observer/display"
)

func main() {
	weather := &observer.Weather{}

	// create displays(Observer)
	currentDisplay := &display.CurrentDisplay{}
	statisticsDisplay := &display.StatisticsDisplay{}

	// add Observers
	weather.AddObserver(currentDisplay)
	weather.AddObserver(statisticsDisplay)

	fmt.Println("=== Initial Measuremants ===")
	weather.SetMeasurements(25.0, 65.5, 1013.1)
	// Output:
	// === Initial Measuremant ===
	// ▼ Current Weather
	// 	Temp :        25.0℃
	// 	Humid:        65.5%
	// 	Press:        1013.1hPa
	// ---------------------------
	// ▼ Statistics Weather
	// 	Data Count:  1
	// 	Avg Temp :   25.0℃
	// 	Avg Humid:   65.5%
	// 	Avg Press:   1013.1hPa
	// ---------------------------

	fmt.Println("=== 2nd Measuremants ===")
	weather.SetMeasurements(22.8, 70.1, 1007.8)
	// Output:
	// === 2nd Measuremants ===
	// ▼ Current Weather
	//   Temp :        22.8℃
	//   Humid:        70.1%
	//   Press:        1007.8hPa
	// ---------------------------
	// ▼ Statistics Weather
	//   Data Count:  2
	//   Avg Temp :   23.9℃
	//   Avg Humid:   67.8%
	//   Avg Press:   1010.5hPa
	// ---------------------------
}
