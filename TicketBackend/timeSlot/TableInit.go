package timeSlot

import (
	"TicketBackend/sql"
	"fmt"
)

func TableInit() {
	err := sql.DB.AutoMigrate(&TimeSlot{})
	if err != nil {
		panic("Table Initiation Failed, Error=" + err.Error())
	}

	var count int64
	sql.DB.Model(&TimeSlot{}).Count(&count)

	if count != 0 {
		return
	}

	timeSlot := TimeSlot{
		TimeSlot: startTime,
		Count:    0,
	}

	for timeSlot.TimeSlot.Before(endTime) {
		e := sql.DB.Create(&timeSlot).Error
		fmt.Println(timeSlot.TimeSlot)
		if e != nil {
			panic("TimeSlot slot create Failed, Error=" + e.Error())
		}
		timeSlot.TimeSlot = timeSlot.TimeSlot.Add(timePerSlot)
	}
}
