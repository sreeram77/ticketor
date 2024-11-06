package handlers

import (
	"context"

	"ticketor/models"
	protogen "ticketor/protogen/proto"
	"ticketor/store"
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
	user, err := t.Users.Get(request.GetUserId())
	if err != nil {
		return nil, err
	}

	section, seat, err := t.Sections.AllocateSeat()
	if err != nil {
		return nil, err
	}

	created, err := t.Tickets.Create(models.Ticket{
		From:    request.GetFrom(),
		To:      request.GetTo(),
		UserID:  request.GetUserId(),
		Section: section,
		Number:  seat,
	})
	if err != nil {
		return nil, err
	}

	return &protogen.TicketReply{
		Id:     created.ID,
		UserId: user.ID,
		User: &protogen.User{
			Id:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
		Number:  created.Number,
		Section: created.Number,
		From:    created.From,
		To:      created.To,
		Price:   created.Price.String(),
	}, nil
}

// GetTicket fetches a ticket by ID.
func (t *Ticketor) GetTicket(ctx context.Context, request *protogen.TicketIDRequest) (*protogen.TicketReply, error) {
	ticket, err := t.Tickets.Get(request.GetId())
	if err != nil {
		return nil, err
	}

	user, err := t.Users.Get(ticket.UserID)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	err = t.Sections.DeallocateSeat(ticket.Section, ticket.Number)
	if err != nil {
		return nil, err
	}

	err = t.Tickets.Remove(request.GetId())
	if err != nil {
		return nil, err
	}

	return &protogen.Empty{}, nil
}

// ModifyTicket modifies a ticket.
func (t *Ticketor) ModifyTicket(ctx context.Context, request *protogen.TicketRequest) (*protogen.TicketReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Ticketor) mustEmbedUnimplementedTicketorServer() {
	//TODO implement me
	panic("implement me")
}
