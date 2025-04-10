package main

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/bridge"
)

func main() {
	tv := &bridge.TV{}
	radio := &bridge.Radio{}

	tvRemote := bridge.NewRemoteController(tv)
	radioRemote := bridge.NewRemoteController(radio)

	fmt.Println(tvRemote.TurnOn())
	fmt.Println(tvRemote.TurnOff())
	// Output:
	// TV is turned on
	// TV is turned off

	fmt.Println(radioRemote.TurnOn())
	fmt.Println(radioRemote.TurnOff())
	// Output:
	// Radio is turned on
	// Radio is turned off
}
