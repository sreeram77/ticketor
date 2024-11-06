package store

import (
	"testing"

	"ticketor/errors"
	"ticketor/models"

	"github.com/golang/mock/gomock"
)

//go:generate mockgen -destination=./mock_users.go -package=store . Users

func TestUsers_Create(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockUser := NewMockUsers(mockCtl)

	user1 := models.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "aaa@example.com",
	}

	gomock.InOrder(
		mockUser.EXPECT().Create(models.User{}).Return(models.User{}, nil),
		mockUser.EXPECT().Create(user1).Return(user1, nil).Times(1),
	)

	_, err := mockUser.Create(models.User{})
	if err != nil {
		t.Fatal(err)
	}

	created, err := mockUser.Create(user1)
	if err != nil {
		t.Fatal(err)
	}

	if created.Email != user1.Email {
		t.Errorf("expected %v, got %v", user1, created)
	}
}

func TestUsers_Get(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockUser := NewMockUsers(mockCtl)

	user1 := models.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "aaa@example.com",
	}

	gomock.InOrder(
		mockUser.EXPECT().Create(user1).Return(user1, nil).Times(1),
		mockUser.EXPECT().Get("1").Return(user1, nil).Times(1),
	)

	created, err := mockUser.Create(user1)
	if err != nil {
		t.Fatal(err)
	}

	got, err := mockUser.Get(created.ID)
	if err != nil {
		t.Fatal(err)
	}

	if got.Email != user1.Email {
		t.Errorf("expected %v, got %v", user1, got)
	}
}

func TestUsers_Remove(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockUser := NewMockUsers(mockCtl)

	user1 := models.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "aaa@example.com",
	}

	gomock.InOrder(
		mockUser.EXPECT().Create(user1).Return(user1, nil).Times(1),
		mockUser.EXPECT().Remove("1").Return(nil).Times(1),
		mockUser.EXPECT().Get("1").Return(models.User{}, errors.ErrNotFound).Times(1),
	)

	created, err := mockUser.Create(user1)
	if err != nil {
		t.Fatal(err)
	}

	err = mockUser.Remove(created.ID)
	if err != nil {
		t.Fatal(err)
	}

	got, err := mockUser.Get(created.ID)
	if err == nil {
		t.Errorf("expected error %v, got %v", errors.ErrNotFound, got)
	}
}
