package display

import "fmt"

// CurrentDisplay displays current mesurements(Observer)
type CurrentDisplay struct {
	temp     float64
	humidity float64
	pressure float64
}

func (cd *CurrentDisplay) Update(temp, humidity, pressure float64) {
	cd.temp = temp
	cd.humidity = humidity
	cd.pressure = pressure
	cd.Display()
}

func (cd *CurrentDisplay) Display() {
	fmt.Println("▼ Current Weather")
	fmt.Printf("  Temp :	%.1f℃\n", cd.temp)
	fmt.Printf("  Humid:	%.1f%%\n", cd.humidity)
	fmt.Printf("  Press:	%.1fhPa\n", cd.pressure)
	fmt.Println("---------------------------")
}
