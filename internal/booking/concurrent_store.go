package booking

import "sync"

type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrencyStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

func (ms *ConcurrentStore) Book(b Booking) error {
	ms.Lock()
	defer ms.Unlock()
	if _, exists := ms.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyBooked
	}
	ms.bookings[b.SeatID] = b

	return nil

}

func (ms *ConcurrentStore) ListBookings(movieID string) []Booking {
	ms.RLock()
	defer ms.RUnlock()
	var result []Booking

	for _, b := range ms.bookings {

		if b.MovieID == movieID {
			result = append(result, b)
		}

	}
	return result

}
