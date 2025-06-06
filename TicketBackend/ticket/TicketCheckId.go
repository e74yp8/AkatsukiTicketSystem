package ticket

import (
	"TicketBackend/sql"
	"fmt"
)

func CheckId(ticketCode string) (id uint, err error) {
	if len(ticketCode) != codeLength {
		return 0, fmt.Errorf("ticketcode length not match, given length: %v, required length: %v", len(ticketCode), codeLength)
	}

	var ticketList []Ticket
	sql.DB.Find(&ticketList, "ticket_code = ?", ticketCode)

	return ticketList[0].Id, nil
}
