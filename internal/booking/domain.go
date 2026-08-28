package booking

import "errors"

// Booking represents a confirmed seat reservation.
type Booking struct {
	ID      string
	MovieID string
	SeatID  string
	UserID  string
	Status  string
}

type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}

var (
	ErrSeatAlreadyBooked = errors.New("Seat Already booked!")
)
