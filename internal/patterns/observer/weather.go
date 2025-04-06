package observer

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
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *Weather) NotifyObservers() {
	for _, obserber := range w.observers {
		obserber.Update(w.temp, w.humidity, w.pressure)
	}
}
