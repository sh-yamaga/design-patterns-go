package singleton

import "sync"

type TicketManager struct {
	currentTicket uint
}

var ticketManager *TicketManager
var once sync.Once

// New returns a new instance of TicketManager only once.
func New() *TicketManager {
	once.Do(func() {
		ticketManager = &TicketManager{currentTicket: 0}
	})
	return ticketManager
}

func (tm *TicketManager) Next() uint {
	tm.currentTicket++
	return tm.currentTicket
}
