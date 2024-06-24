package ticket

import (
	"TicketBackend/sql"
	"fmt"
)

func CheckExist(ticketCode string) (contain bool, err error) {
	if len(ticketCode) != codeLength {
		return false, fmt.Errorf("ticketcode length not match, given length: %v, required length: %v", len(ticketCode), codeLength)
	}

	var ticketList []Ticket
	sql.DB.Find(&ticketList, "ticket_code = ?", ticketCode)
	if len(ticketList) == 0 {
		return false, nil
	}
	return true, nil
}
