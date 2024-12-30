package ticket

import (
	"TicketBackend/sql"
	"fmt"
)

// UpdateStatus
//
//		Status Code:
//	  -  0: Not used
//	  -  1: Used
//	  - -1: Disabled
func UpdateStatus(ticketCode string, status int) (err error) {
	if len(ticketCode) != codeLength {
		return fmt.Errorf("ticketcode length not match, given length: %v, required length: %v", len(ticketCode), codeLength)
	}

	var ticket Ticket
	sql.DB.Take(&ticket, "ticket_code = ?", ticketCode)
	ticket.Status = status
	sql.DB.Select("status").Save(&ticket)

	return nil
}
