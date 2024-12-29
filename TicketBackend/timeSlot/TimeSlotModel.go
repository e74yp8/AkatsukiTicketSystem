package timeSlot

import "time"

type TimeSlot struct {
	TimeSlot time.Time `gorm:"primaryKey;unique;not null"`
	Count    uint      `gorm:"default:0"`
}

var loc, _ = time.LoadLocation("Asia/Hong_Kong")
var startTime = time.Date(2025, 1, 19, 12, 0, 0, 0, loc)
var endTime = time.Date(2025, 1, 19, 17, 30, 0, 0, loc)
var timePerSlot = time.Minute * 30

var LimitPerTimeSlot uint = 150
var FirstLimitPerTimeSlot uint = 300
