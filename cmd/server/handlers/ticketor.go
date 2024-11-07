package handlers

import (
	"context"

	"ticketor/models"
	protogen "ticketor/protogen/proto"
	"ticketor/store"
	"ticketor/utils"
)

type Ticketor struct {
	Users    store.Users
	Sections store.Sections
	Tickets  store.Tickets

	protogen.UnimplementedTicketorServer
}

// NewTicketor creates a new ticketor server.
func NewTicketor(users store.Users, sections store.Sections, tickets store.Tickets) protogen.TicketorServer {
	return &Ticketor{
		Users:    users,
		Sections: sections,
		Tickets:  tickets,
	}
}

// PurchaseTicket creates a new ticket.
func (t *Ticketor) PurchaseTicket(ctx context.Context, request *protogen.TicketRequest) (*protogen.TicketReply, error) {
	fetchedUser, err := t.Users.Get(request.GetUserId())
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	section, seat, err := t.Sections.AllocateSeat()
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	created, err := t.Tickets.Create(models.Ticket{
		From:    request.GetFrom(),
		To:      request.GetTo(),
		UserID:  request.GetUserId(),
		Section: section,
		Number:  seat,
		Price:   models.NewMoney(20),
	})
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	return &protogen.TicketReply{
		Id:     created.ID,
		UserId: fetchedUser.ID,
		User: &protogen.User{
			Id:        fetchedUser.ID,
			FirstName: fetchedUser.FirstName,
			LastName:  fetchedUser.LastName,
			Email:     fetchedUser.Email,
		},
		Number:  created.Number,
		Section: created.Section,
		From:    created.From,
		To:      created.To,
		Price:   created.Price.String(),
	}, nil
}

// GetTicket fetches a ticket by ID.
func (t *Ticketor) GetTicket(ctx context.Context, request *protogen.TicketIDRequest) (*protogen.TicketReply, error) {
	ticket, err := t.Tickets.Get(request.GetId())
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	user, err := t.Users.Get(ticket.UserID)
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	return &protogen.TicketReply{
		Id:      ticket.ID,
		UserId:  user.ID,
		Number:  ticket.Number,
		Section: ticket.Section,
		From:    ticket.From,
		To:      ticket.To,
		Price:   ticket.Price.String(),
	}, nil
}

// RemoveTicket deletes a ticket.
func (t *Ticketor) RemoveTicket(ctx context.Context, request *protogen.TicketIDRequest) (*protogen.Empty, error) {
	ticket, err := t.Tickets.Get(request.GetId())
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	err = t.Sections.DeallocateSeat(ticket.Section, ticket.Number)
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	err = t.Tickets.Remove(request.GetId())
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	return &protogen.Empty{}, nil
}

// ModifyTicket modifies a ticket.
func (t *Ticketor) ModifyTicket(ctx context.Context, request *protogen.TicketRequest) (*protogen.TicketReply, error) {
	ticket, err := t.Tickets.Get(request.GetId())
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	// Deallocate existing ticket.
	err = t.Sections.DeallocateSeat(ticket.Section, ticket.Number)
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	// Allocate new seat.
	section, seat, err := t.Sections.AllocateSeat()
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	modified, err := t.Tickets.Modify(request.GetId(), models.Ticket{
		ID:      request.GetId(),
		From:    request.GetFrom(),
		To:      request.GetTo(),
		UserID:  ticket.UserID,
		Section: section,
		Number:  seat,
	})
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	return &protogen.TicketReply{
		Id:      modified.ID,
		UserId:  modified.UserID,
		From:    modified.From,
		To:      modified.To,
		Section: modified.Section,
		Number:  modified.Number,
		Price:   modified.Price.String(),
		User: &protogen.User{
			Id:        modified.User.ID,
			FirstName: modified.User.FirstName,
			LastName:  modified.User.LastName,
			Email:     modified.User.Email,
		},
	}, nil
}

func (t *Ticketor) mustEmbedUnimplementedTicketorServer() {}
