package entry

import (
	"TicketBackend/googlesheet"
	"TicketBackend/ticket"
	"fmt"
)

func TicketCheck(ticketCode string) (err error) {
	exist, err := ticket.CheckExist(ticketCode)
	if err != nil || !exist {
		return fmt.Errorf("門票不存在")
	}
	status, err := ticket.CheckStatus(ticketCode)
	if err != nil {
		return err
	} else if status != 0 {
		if status == 1 {
			return fmt.Errorf("門票已被使用")
		} else if status == -1 {
			return fmt.Errorf("門票已被取消")
		}
	}

	id, err := ticket.CheckId(ticketCode)
	if err != nil {
		return err
	}

	if id > uint(googlesheet.GetLimitId()) {
		return fmt.Errorf("未到籌號")
	}

	err = ticket.UpdateStatus(ticketCode, 1)
	if err != nil {
		return err
	}
	return nil
}
