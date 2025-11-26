package handlers

import (
	"awesomeProject/internal/userService"
	"awesomeProject/internal/web/users"
	"context"
)

type UserHandler struct {
	Service userService.UserService
}

func NewUserHandler(s userService.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}
	for _, u := range allUsers {
		user := users.User{
			Email:    &u.Email,
			Id:       &u.ID,
			Password: &u.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id

	user := request.Body

	updatedUser, err := h.Service.UpdateUser(id, &userService.User{
		Email:    *user.Email,
		Password: *user.Password,
	})

	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToRequest := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(&userToRequest)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Email:    &createdUser.Email,
		Id:       &createdUser.ID,
		Password: &createdUser.Password,
	}
	return response, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id
	err := h.Service.DeleteUser(id)
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204JSONResponse{}, nil
}
