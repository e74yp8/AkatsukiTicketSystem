package ticket

import (
	"TicketBackend/sql"
	"fmt"
)

func Create(ticketCode string, email string) (id uint, err error) {
	ticket1 := Ticket{
		TicketCode: ticketCode,
		Email:      email,
	}

	e := sql.DB.Create(&ticket1).Error
	if e != nil {
		return 0, fmt.Errorf("ticket Create Failed, Error=" + e.Error())
	}

	id, e = CheckId(ticketCode)
	if e != nil {
		return 0, fmt.Errorf("ticket Create Failed, Error=" + e.Error())
	}

	return id, nil
}
