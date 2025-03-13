package singleton

import "sync"

type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

// New returns a new instance of Singleton
func New() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "initial value"}
	})
	return instance
}
