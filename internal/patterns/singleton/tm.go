package singleton

import "sync"

type TicketManager struct {
	currentTicket uint
}

var ticketManager *TicketManager
var once sync.Once

// NewTicketManager returns a single instance of TicketManager.
func NewTicketManager() *TicketManager {
	once.Do(func() {
		ticketManager = &TicketManager{currentTicket: 0}
	})
	return ticketManager
}

func (tm *TicketManager) Next() uint {
	tm.currentTicket++
	return tm.currentTicket
}
