package ticket

import (
	"TicketBackend/sql"
	"TicketBackend/timeSlot"
	"fmt"
	"time"
)

func Create(ticketCode string, slot time.Time) (err error) {
	ticket1 := Ticket{
		TicketCode: ticketCode,
		TimeSlot:   slot,
	}

	e := sql.DB.Create(&ticket1).Error
	if e != nil {
		return fmt.Errorf("ticket Create Failed, Error=" + e.Error())
	}

	timeSlot.CountUpdate(slot)

	return nil
}
