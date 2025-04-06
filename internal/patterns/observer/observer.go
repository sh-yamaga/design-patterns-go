package observer

// Observer is responsible for receiving updates from the Subject and reacting to changes in state.
type Observer interface {
	Update(temp, humidity, pressure float64)
}
