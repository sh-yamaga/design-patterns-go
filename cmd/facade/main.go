package main

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/facade"

func main() {
	travelFacade := facade.NewTravelFacade()

	travelFacade.BookTravel("Tokyo")
}
