package observer

import "slices"

// Weather holds measurement values and notifies Observers
type Weather struct {
	observers []Observer
	temp      float64
	humidity  float64
	pressure  float64
}

func (w *Weather) AddObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *Weather) RemoveObserver(o Observer) {
	for i, observer := range w.observers {
		if observer == o {
			w.observers = slices.Delete(w.observers, i, i+1)
			break
		}
	}
}

func (w *Weather) NotifyObservers() {
	for _, obserber := range w.observers {
		obserber.Update(w.temp, w.humidity, w.pressure)
	}
}

func (w *Weather) SetMeasurements(temp, humidity, pressure float64) {
	w.temp = temp
	w.humidity = humidity
	w.pressure = pressure

	w.NotifyObservers()
}
