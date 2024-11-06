package store

import (
	"errors"

	"ticketor/models"

	"github.com/google/uuid"
)

type tickets struct {
	store map[string]models.Ticket
}

// NewTickets creates a new tickets store.
func NewTickets() Tickets {
	return &tickets{
		store: make(map[string]models.Ticket),
	}
}

// Create creates a new ticket.
func (t *tickets) Create(ticket models.Ticket) (models.Ticket, error) {
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
