package main

import "github.com/yamaga-shu/design-patterns-go/internal/patterns/facade"

func main() {
	travelFacade := facade.NewTravelFacade()

	travelFacade.BookTravel("Tokyo")
	// Output:
	// booked flight: destination -> Tokyo
	// booked hotel: location -> Tokyo
	// booked rental car: location -> Tokyo
}
