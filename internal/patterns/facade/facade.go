package facade

import "github.com/yamaga-shu/design-patterns-go/internal/patterns/facade/book"

type TravelFacade struct {
	flightBooker    *book.FlightBooker
	hotelBooker     *book.HotelBooker
	rentalCarBooker *book.RentalCarBooker
}

func NewTravelFacade() *TravelFacade {
	return &TravelFacade{
		flightBooker:    &book.FlightBooker{},
		hotelBooker:     &book.HotelBooker{},
		rentalCarBooker: &book.RentalCarBooker{},
	}
}

func (tf *TravelFacade) BookTravel(dest string) {
	tf.flightBooker.BookFlight(dest)
	tf.hotelBooker.BookHotel(dest)
	tf.rentalCarBooker.BookCar(dest)
}
