package timeSlot

import (
	"TicketBackend/sql"
	"time"
)

func CountUpdate(t time.Time) {
	var timeSlot TimeSlot
	sql.DB.Take(&timeSlot, "time_slot = ?", t)
	timeSlot.Count += 1
	sql.DB.Select("count").Save(&timeSlot)
}
