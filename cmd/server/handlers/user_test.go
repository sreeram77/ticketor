package handlers

import (
	"context"
	"testing"

	"ticketor/models"
	protogen "ticketor/protogen/proto"
	"ticketor/store"

	"github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	testUserRequest := &protogen.UserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
	}

	testUserModel := models.User{
		FirstName: testUserRequest.FirstName,
		LastName:  testUserRequest.LastName,
		Email:     testUserRequest.Email,
	}

	mockUserStore := store.NewMockUsers(controller)

	mockUserStore.EXPECT().Create(testUserModel).Return(testUserModel, nil)

	handler := NewUser(mockUserStore)

	created, err := handler.Create(context.Background(), testUserRequest)
	if err != nil {
		return
	}

	if created.Email != testUserRequest.Email {
		t.Errorf("expected %v, got %v", testUserRequest, created)
	}
}

func TestGetUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	testUserRequest := &protogen.UserIDRequest{
		Id: "1",
	}

	testUserModel := models.User{
		ID:        testUserRequest.Id,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
	}

	mockUserStore := store.NewMockUsers(controller)

	mockUserStore.EXPECT().Get(testUserRequest.Id).Return(testUserModel, nil)

	handler := NewUser(mockUserStore)

	created, err := handler.Get(context.Background(), testUserRequest)
	if err != nil {
		return
	}

	if created.Id != testUserRequest.GetId() {
		t.Errorf("expected %v, got %v", testUserRequest, created)
	}
}
