package ticket

import "TicketBackend/sql"

func TableInit() {
	err := sql.DB.AutoMigrate(&Ticket{})
	if err != nil {
		panic("Table Initiation Failed, Error=" + err.Error())
	}
}
