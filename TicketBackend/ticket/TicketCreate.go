package ticket

import "TicketBackend/sql"

func Create(ticketCode string) {
	ticket1 := Ticket{
		TicketCode: ticketCode,
	}

	err := sql.DB.Create(&ticket1).Error
	if err != nil {
		panic("ticket Create Failed, Error=" + err.Error())
	}

}
