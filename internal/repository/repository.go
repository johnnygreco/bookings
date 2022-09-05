package repository

import (
	"github.com/johnnygreco/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error) 
	InsertRoomRestriction(r models.RoomRestriction) error
}
