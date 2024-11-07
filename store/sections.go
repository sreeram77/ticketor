package store

import (
	"strconv"

	"ticketor/errors"
	"ticketor/models"
)

const (
	maxSeats = 61
)

type sections struct {
	store map[string]models.Section
}

func NewSections() Sections {
	store := make(map[string]models.Section)

	// Populate data.
	store["1"] = models.Section{ID: "1", Number: 1, Occupancy: make([]bool, maxSeats)}
	store["2"] = models.Section{ID: "2", Number: 2, Occupancy: make([]bool, maxSeats)}

	return &sections{
		store: store,
	}
}

// Get fetches a section by ID.
func (s *sections) Get(id string) (models.Section, error) {
	section, exists := s.store[id]
	if !exists {
		return models.Section{}, errors.ErrNotFound
	}

	return section, nil
}

// AllocateSeat creates a new seat.
func (s *sections) AllocateSeat() (string, string, error) {
	for k, v := range s.store {
		section := v.Occupancy
		for j := range section {
			// 0 is ignored
			if j == 0 {
				continue
			}

			if !section[j] {
				section[j] = true

				return k, strconv.Itoa(j), nil
			}
		}
	}

	return "", "", errors.ErrNotAvailable
}

func (s *sections) DeallocateSeat(section, seat string) error {
	sec, err := s.Get(section)
	if err != nil {
		return err
	}

	seatNumber, err := strconv.Atoi(seat)
	if err != nil {
		return errors.ErrInvalid
	}

	// Check if seat is occupied.
	occupied := sec.Occupancy[seatNumber]
	if !occupied {
		return errors.ErrInvalid
	}

	// Deallocate seat.
	sec.Occupancy[seatNumber] = false

	return nil
}

func (s *sections) ReallocateSeat(section, seat string) (string, string, error) {
	for k, v := range s.store {
		sec := v.Occupancy
		for j := range sec {
			// 0 is ignored
			if j == 0 {
				continue
			}

			if !sec[j] {
				sec[j] = true

				// deallocate older seat.
				err := s.DeallocateSeat(section, seat)
				if err != nil {
					return "", "", err
				}

				return k, strconv.Itoa(j), nil
			}
		}
	}

	return "", "", errors.ErrNotAvailable
}
