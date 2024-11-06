package store

import (
	"ticketor/errors"
	"ticketor/models"

	"github.com/google/uuid"
)

type users struct {
	store map[string]models.User
}

// NewUsers creates a new users store.
func NewUsers() Users {
	return &users{
		store: make(map[string]models.User),
	}
}

// Create creates a new user.
func (u *users) Create(user models.User) (models.User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return models.User{}, err
	}

	user.ID = id.String()

	u.store[user.ID] = user

	return user, nil

}

// Get fetches a user by ID.
func (u *users) Get(id string) (models.User, error) {
	user, ok := u.store[id]
	if !ok {
		return models.User{}, errors.ErrNotFound
	}

	return user, nil
}

// Remove deletes a user by ID.
func (u *users) Remove(id string) error {
	user, err := u.Get(id)
	if err != nil {
		return err
	}

	delete(u.store, user.ID)

	return nil
}
