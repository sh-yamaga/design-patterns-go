package book

import "fmt"

type FlightBooker struct{}

func (fb *FlightBooker) BookFlight(dest string) {
	fmt.Printf("booked flight: destination -> %s\n", dest)
}
