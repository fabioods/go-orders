package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/fabioods/go-orders/internal/errorcode"
	"github.com/fabioods/go-orders/internal/infra/webserver"
	"github.com/fabioods/go-orders/internal/usecase"
	"github.com/fabioods/go-orders/pkg/errorformatted"
	"github.com/fabioods/go-orders/pkg/response"
	"github.com/fabioods/go-orders/pkg/trace"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	CreateUserUseCase CreateUserUseCase
	UserAvatarUseCase UserAvatarUseCase
}

//go:generate mockery --name=CreateUserUseCase --output=mocks --case=underscore
type CreateUserUseCase interface {
	Execute(ctx context.Context, input usecase.CreateUserDTO) error
}

//go:generate mockery --name=UserAvatarUseCase --output=mocks --case=underscore
type UserAvatarUseCase interface {
	Execute(ctx context.Context, dto usecase.UserAvatarDTO) error
}

func NewUserHandler(createUserUseCase CreateUserUseCase) *UserHandler {
	return &UserHandler{
		CreateUserUseCase: createUserUseCase,
	}
}

func (h *UserHandler) AddUserHandler(web *webserver.WebServer) {
	web.AddRoute(http.MethodPost, "/users", h.addUser)
	web.AddRoute(http.MethodPost, "/users/avatar/{userId}", h.addUserAvatar)
}

func (h *UserHandler) addUser(w http.ResponseWriter, r *http.Request) {
	var userDto usecase.CreateUserDTO
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		ef := errorformatted.BadRequestError(trace.GetTrace(), errorcode.ErrorUserDecodeError, "%s", err.Error())
		response.WriteResponse(w, nil, ef, http.StatusBadRequest)
	}

	err = h.CreateUserUseCase.Execute(r.Context(), usecase.CreateUserDTO{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	})

	response.WriteResponse(w, nil, err, http.StatusCreated)
}

func (h *UserHandler) addUserAvatar(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		ef := errorformatted.BadRequestError(trace.GetTrace(), errorcode.ErrorAvatarFileError, "%s", err.Error())
		response.WriteResponse(w, nil, ef, http.StatusBadRequest)
	}

	defer file.Close()

	userId := chi.URLParam(r, "userId")
	if userId == "" {
		ef := errorformatted.BadRequestError(trace.GetTrace(), errorcode.ErrorUserIdNotProvided, "%s", err.Error())
		response.WriteResponse(w, nil, ef, http.StatusBadRequest)
	}

	err = h.UserAvatarUseCase.Execute(r.Context(), usecase.UserAvatarDTO{
		UserID: userId,
		Avatar: file,
	})

	response.WriteResponse(w, nil, err, http.StatusCreated)

}
