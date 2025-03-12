package usecase

import (
	"context"

	"github.com/fabioods/go-orders/internal/entity"
	"github.com/fabioods/go-orders/internal/errorcode"
	"github.com/fabioods/go-orders/pkg/errorformatted"
	"github.com/fabioods/go-orders/pkg/trace"
)

type (
	CreateUserDTO struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CreateUserUseCase struct {
		UserRepository   UserRepository
		UploadRepository UploadRepository
	}
)

//go:generate mockery --name=UserRepository --output=mocks --case=underscore
type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, id string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
}

func NewCreateUserUseCase(userRepository UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, dto CreateUserDTO) error {
	user := entity.NewUser()
	user.Name = dto.Name
	user.Email = dto.Email

	errPassword := user.SetPassword(dto.Password)
	if errPassword != nil {
		return errPassword
	}

	validateUser := user.Validate()
	if validateUser != nil {
		ef := errorformatted.BadRequestError(trace.GetTrace(), errorcode.ErrorUserValidate, "%s", validateUser.Error())
		return ef
	}

	err := uc.UserRepository.Save(ctx, user)
	if err != nil {
		return err
	}

	return nil

}
