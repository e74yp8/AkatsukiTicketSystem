package ticket

import "time"

type Ticket struct {
	Id         uint   `gorm:"primaryKey;autoIncrement;unique;not null"`
	TicketCode string `gorm:"size:128"`
	TimeSlot   time.Time
	Status     int    `gorm:"size:1;default:0"`
	Email      string `gorm:"size:64"` // TODO: ticket要綁定email
	// TODO: resend qrcode function
}

var codeLength = 128
