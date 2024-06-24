package ticket

type Ticket struct {
	Id         uint   `gorm:"primaryKey;autoIncrement;unique;not null"`
	TicketCode string `gorm:"size:128"`
	Status     int    `gorm:"size:1;default:0"`
}

var codeLength = 10
