package entry

import (
	"TicketBackend/sql"
	"TicketBackend/timeSlot"
	"time"
)

func TimeSelect() (time time.Time) {
	var slotList []timeSlot.TimeSlot
	sqlStatement := "SELECT time_slot FROM time_slots WHERE ((DATETIME(time_slot) BETWEEN '2025-01-19 11:59:59' AND '2025-01-19 12:00:01') AND (count < ?)) OR (count < ?)"
	sql.DB.Raw(sqlStatement, timeSlot.FirstLimitPerTimeSlot, timeSlot.LimitPerTimeSlot).Scan(&slotList)

	//for _, slot := range slotList {
	//	fmt.Println(slot.TimeSlot)
	//}

	return slotList[0].TimeSlot
}
