package handlers

import (
	"context"

	"ticketor/errors"
	"ticketor/models"
	protogen "ticketor/protogen/proto"
	"ticketor/store"
	"ticketor/utils"
)

type user struct {
	store store.Users

	protogen.UnimplementedUserServiceServer
}

// NewUser creates a new user handler.
func NewUser(users store.Users) protogen.UserServiceServer {
	return &user{
		store: users,
	}
}

// Get fetches a user by ID.
func (u *user) Get(ctx context.Context, request *protogen.UserIDRequest) (*protogen.UserResponse, error) {
	fetched, err := u.store.Get(request.GetId())
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	return &protogen.UserResponse{
		Id:        fetched.ID,
		FirstName: fetched.FirstName,
		LastName:  fetched.LastName,
		Email:     fetched.Email,
	}, nil
}

// Create creates a new user.
func (u *user) Create(ctx context.Context, request *protogen.UserRequest) (*protogen.UserResponse, error) {
	err := validateUser(request)
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	created, err := u.store.Create(models.User{
		FirstName: request.GetFirstName(),
		LastName:  request.GetLastName(),
		Email:     request.GetEmail(),
	})
	if err != nil {
		return nil, utils.StatusFromError(err)
	}

	return &protogen.UserResponse{
		Id:        created.ID,
		FirstName: created.FirstName,
		LastName:  created.LastName,
		Email:     created.Email,
	}, nil
}

func validateUser(request *protogen.UserRequest) error {
	if request.GetEmail() == "" || request.GetFirstName() == "" || request.GetLastName() == "" {
		return errors.ErrInvalid
	}

	if !utils.CheckEmail(request.GetEmail()) {
		return errors.ErrInvalid
	}

	return nil
}

func (u *user) mustEmbedUnimplementedUserServer() {}
