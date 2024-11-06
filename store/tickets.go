package store

import (
	"ticketor/errors"
	"ticketor/models"

	"github.com/google/uuid"
)

type tickets struct {
	store        map[string]models.Ticket
	destinations map[string]struct{}
}

// NewTickets creates a new tickets store.
func NewTickets() Tickets {
	destinations := make(map[string]struct{})

	// Populate data
	destinations["London"] = struct{}{}
	destinations["France"] = struct{}{}

	return &tickets{
		store:        make(map[string]models.Ticket),
		destinations: destinations,
	}
}

// Create creates a new ticket.
func (t *tickets) Create(ticket models.Ticket) (models.Ticket, error) {
	if !t.checkDestinationExists(ticket.To) || !t.checkDestinationExists(ticket.From) {
		return models.Ticket{}, errors.ErrNotFound
	}

	if ticket.From == ticket.To {
		return models.Ticket{}, errors.ErrInvalid
	}

	id, err := uuid.NewV7()
	if err != nil {
		return models.Ticket{}, err
	}

	ticket.ID = id.String()
	t.store[ticket.ID] = ticket

	return ticket, nil
}

// Get fetches a ticket by ID.
func (t *tickets) Get(id string) (models.Ticket, error) {
	ticket, exists := t.store[id]
	if !exists {
		return models.Ticket{}, errors.ErrNotFound
	}

	return ticket, nil
}

// Remove deletes a ticket by ID.
func (t *tickets) Remove(id string) error {
	ticket, err := t.Get(id)
	if err != nil {
		return err
	}

	delete(t.store, ticket.ID)

	return nil
}

// Modify modifies a ticket.
func (t *tickets) Modify(id string, ticket models.Ticket) (models.Ticket, error) {
	ticket.ID = id
	t.store[id] = ticket

	return ticket, nil
}

// GetBySection fetches all tickets by section ID.
func (t *tickets) GetBySection(id string) ([]models.Ticket, error) {
	var res []models.Ticket

	for _, v := range t.store {
		if v.Section == id {
			res = append(res, v)
		}
	}

	return res, nil
}

// checkDestinationExists checks if a destination exists.
func (t *tickets) checkDestinationExists(destination string) bool {
	_, exists := t.destinations[destination]

	return exists
}
