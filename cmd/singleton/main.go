package main

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/singleton"
)

func main() {
	tm1 := singleton.NewTicketManager()
	tm2 := singleton.NewTicketManager()

	fmt.Println("Next ticketNo:", tm1.Next())
	fmt.Println("Next ticketNo:", tm2.Next())
	fmt.Println("Next ticketNo:", tm1.Next())
	fmt.Println("Next ticketNo:", tm2.Next())
	fmt.Println("Next ticketNo:", tm1.Next())
	// Output:
	// Next ticketNo: 1
	// Next ticketNo: 2
	// Next ticketNo: 3
	// Next ticketNo: 4
	// Next ticketNo: 5

	if tm1 == tm2 {
		fmt.Println("tm1 == tm2")
	}
	// Output:
	// tm1 == tm2
}
