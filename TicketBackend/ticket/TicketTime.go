package ticket

import (
	"TicketBackend/sql"
	"time"
)

func CheckTime(ticketCode string) (valid bool) {
	if len(ticketCode) != codeLength {
		panic("Ticket code length does not match.")
	}

	var ticketList []Ticket
	sql.DB.Find(&ticketList, "ticket_code = ?", ticketCode)

	return ticketList[0].TimeSlot.Before(time.Now())
}
