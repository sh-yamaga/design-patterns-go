package book

import "fmt"

type HotelBooker struct{}

func (hb *HotelBooker) BookHotel(location string) {
	fmt.Printf("booked hotel: location -> %s\n", location)
}
