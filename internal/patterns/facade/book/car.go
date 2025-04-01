package book

import "fmt"

type RentalCarBooker struct{}

func (rcb *RentalCarBooker) BookCar(location string) {
	fmt.Printf("booked rental car: location -> %s\n", location)
}
