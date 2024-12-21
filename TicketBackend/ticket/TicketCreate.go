package ticket

import (
	"TicketBackend/sql"
	"fmt"
)

func Create(ticketCode string) (err error) {
	ticket1 := Ticket{
		TicketCode: ticketCode,
	}

	e := sql.DB.Create(&ticket1).Error
	if e != nil {
		return fmt.Errorf("ticket Create Failed, Error=" + e.Error())
	}
	return nil
}
