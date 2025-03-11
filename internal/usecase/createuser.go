package usecase

import (
	"context"
	"mime/multipart"

	"github.com/fabioods/go-orders/internal/entity"
	"github.com/fabioods/go-orders/internal/errorcode"
	"github.com/fabioods/go-orders/pkg/errorformatted"
	"github.com/fabioods/go-orders/pkg/rollback"
	"github.com/fabioods/go-orders/pkg/trace"
)

type (
	CreateUserDTO struct {
		Name     string         `json:"name"`
		Email    string         `json:"email"`
		Password string         `json:"password"`
		Avatar   multipart.File `json:"avatar"`
	}

	CreateUserUseCase struct {
		UserRepository   UserRepository
		UploadRepository UploadRepository
	}

	UploadRepository interface {
		Upload(ctx context.Context, file multipart.File, fileName string) (string, error)
		Delete(ctx context.Context, fileName string) error
	}

	UserRepository interface {
		Save(ctx context.Context, user *entity.User) error
		FindByID(ctx context.Context, id string) (*entity.User, error)
	}
)

func NewCreateUserUseCase(userRepository UserRepository, uploadRepository UploadRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository:   userRepository,
		UploadRepository: uploadRepository,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, dto CreateUserDTO) error {
	rb := rollback.New()

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

	avatarURL, err := uc.UploadRepository.Upload(ctx, dto.Avatar, user.ID)
	if err != nil {
		return err
	}
	user.SetAvatar(avatarURL)
	rb.Add("Avatar upload", func() error {
		err := uc.UploadRepository.Delete(ctx, user.ID)
		if err != nil {
			return err
		}
		return nil
	})

	err = uc.UserRepository.Save(ctx, user)
	if err != nil {
		rb.Do(ctx)
		return err
	}

	return nil

}
