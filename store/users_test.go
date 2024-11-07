package store

import (
	"errors"
	"testing"

	intErrors "ticketor/errors"
	"ticketor/models"
)

//go:generate mockgen -destination=./mock_users.go -package=store . Users

func TestUsers_Create(t *testing.T) {
	u := NewUsers()

	testUser := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
	}

	got, err := u.Create(testUser)
	if err != nil {
		t.Fatal(err)
	}

	if got.Email != testUser.Email {
		t.Errorf("expected %v, got %v", testUser, got)
	}
}

func TestUsers_Get(t *testing.T) {
	u := NewUsers()

	testUser := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
	}

	created, err := u.Create(testUser)
	if err != nil {
		t.Fatal(err)
	}

	got, err := u.Get(created.ID)
	if err != nil {
		t.Fatal(err)
	}

	if got.Email != testUser.Email {
		t.Errorf("expected %v, got %v", testUser, got)
	}
}

func TestUsers_Remove(t *testing.T) {
	u := NewUsers()

	testUser := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
	}

	created, err := u.Create(testUser)
	if err != nil {
		t.Fatal(err)
	}

	err = u.Remove(created.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = u.Get(created.ID)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if !errors.Is(err, intErrors.ErrNotFound) {
		t.Errorf("expected %v, got %v", intErrors.ErrNotFound, err)
	}
}
