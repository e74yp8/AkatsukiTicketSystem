package entry

import (
	"TicketBackend/ticket"
	"fmt"
)

func TicketCheck(ticketCode string) (err error) {
	exist, err := ticket.CheckExist(ticketCode)
	if err != nil {
		return err
	} else if !exist {
		return fmt.Errorf("ticket not exist, ticketcode: %v", ticketCode)
	}
	status, err := ticket.CheckStatus(ticketCode)
	if err != nil {
		return err
	} else if status != 0 {
		if status == 1 {
			return fmt.Errorf("ticket has been used, ticketcode: %v", ticketCode)
		} else if status == -1 {
			return fmt.Errorf("ticket has been disabled, ticketcode: %v", ticketCode)
		}
	}
	err = ticket.UpdateStatus(ticketCode, 1)
	if err != nil {
		return err
	}
	return nil
}
