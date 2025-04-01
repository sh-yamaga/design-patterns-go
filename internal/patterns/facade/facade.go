package facade

type TravelFacade struct {
	flightBooker    *book.FlightBooker
	hotelBooker     *book.HotelBooker
	rentalCarBooker *book.RentalCarBooker
}

func NewTravelFacade() *TravelTravelFacade {
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
