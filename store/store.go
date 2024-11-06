package store

import "ticketor/models"

type Users interface {
	Create(user models.User) (models.User, error)
	Get(id string) (models.User, error)
	Remove(id string) error
}

type Tickets interface {
	Create(ticket models.Ticket) (models.Ticket, error)
	Get(id string) (models.Ticket, error)
	GetBySection(id string) ([]models.Ticket, error)
	Remove(id string) error
	Modify(id string, ticket models.Ticket) (models.Ticket, error)
}

type Sections interface {
	Get(id string) (models.Section, error)
	AllocateSeat() (string, string, error)
	DeallocateSeat(section, seat string) error
}
