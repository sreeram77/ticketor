package handlers

import (
	"context"

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

func (t *Ticketor) PurchaseTicket(ctx context.Context, request *protogen.TicketRequest) (*protogen.TicketReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Ticketor) GetTicket(ctx context.Context, request *protogen.TicketIDRequest) (*protogen.TicketReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Ticketor) RemoveUser(ctx context.Context, request *protogen.TicketRequest) (*protogen.TicketReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Ticketor) ModifyTicket(ctx context.Context, request *protogen.TicketRequest) (*protogen.TicketReply, error) {
	//TODO implement me
	panic("implement me")
}

func (t Ticketor) mustEmbedUnimplementedTicketorServer() {
	//TODO implement me
	panic("implement me")
}
