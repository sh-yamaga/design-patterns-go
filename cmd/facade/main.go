package main

func main() {
	travelFacade := facade.NewTravelFacade()

	travelFacade.BookTravel("Tokyo")
}
