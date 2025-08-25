// Package entities
package entities

type TicketStatus string

const (
	TicketApproved  TicketStatus = "approved"
	TicketCancelled string       = "cancelled"
)

type Ticket struct {
	ID       uint         `gorm:"primaryKey;autoIncrement"`
	SeatID   uint         `gorm:"not null"`
	Username string       `gorm:"type:varchar(100);not null"`
	Status   TicketStatus `gorm:"type:ticket_status;not null"`

	Seat Seat `gorm:"foreignKey:SeatID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
