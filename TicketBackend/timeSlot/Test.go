package timeSlot

import (
	"TicketBackend/sql"
	"time"
)

func CountSet(t time.Time, count uint) {
	var timeSlot TimeSlot
	sql.DB.Take(&timeSlot, "time_slot = ?", t)
	timeSlot.Count = count
	sql.DB.Select("count").Save(&timeSlot)
}
