package display

import (
	"fmt"
)

type StatisticsDisplay struct {
	tempRecord     []float64
	humidRecord    []float64
	pressureRecord []float64
}

func (sd *StatisticsDisplay) Update(temp, humidity, pressure float64) {
	sd.tempRecord = append(sd.tempRecord, temp)
	sd.humidRecord = append(sd.humidRecord, humidity)
	sd.pressureRecord = append(sd.pressureRecord, pressure)
	sd.Display()
}

func (sd *StatisticsDisplay) Display() {
	averageTemp := getAverageMeasurement(sd.tempRecord)
	averageHumid := getAverageMeasurement(sd.humidRecord)
	averagePress := getAverageMeasurement(sd.pressureRecord)

	fmt.Println("▼ Statistics Weather")
	fmt.Printf("  Data Count:  %d\n", len(sd.tempRecord))
	fmt.Printf("  Avg Temp :   %.1f℃\n", averageTemp)
	fmt.Printf("  Avg Humid:   %.1f%%\n", averageHumid)
	fmt.Printf("  Avg Press:   %.1fhPa\n", averagePress)
	fmt.Println("---------------------------")
}

func getAverageMeasurement(records []float64) float64 {
	var sum float64
	for _, v := range records {
		sum += v
	}

	return sum / float64(len(records))
}
