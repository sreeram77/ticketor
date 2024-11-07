package handlers

import (
	"context"

	protogen "ticketor/protogen/proto"
	"ticketor/store"
)

type user struct {
	store store.Users

	protogen.UnimplementedUserServer
}

func NewUser(users store.Users) protogen.UserServer {
	return &user{
		store: users,
	}
}

func (u user) Get(ctx context.Context, request *protogen.UserRequest) (*protogen.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u user) Create(ctx context.Context, request *protogen.UserRequest) (*protogen.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u user) mustEmbedUnimplementedUserServer() {}
