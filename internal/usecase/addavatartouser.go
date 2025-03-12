package usecase

import (
	"context"
	"mime/multipart"

	"github.com/fabioods/go-orders/pkg/errorformatted"
	"github.com/fabioods/go-orders/pkg/trace"
)

type (
	UserAvatarDTO struct {
		UserID string         `json:"user_id"`
		Avatar multipart.File `json:"avatar"`
	}

	UserAvatarUseCase struct {
		UserRepository   UserRepository
		UploadRepository UploadRepository
	}
)

//go:generate mockery --name=UploadRepository --output=mocks --case=underscore
type UploadRepository interface {
	Upload(ctx context.Context, file multipart.File, fileName string) (string, error)
	Delete(ctx context.Context, fileName string) error
}

func NewUserAvatarUseCase(userRepository UserRepository, uploadRepository UploadRepository) *UserAvatarUseCase {
	return &UserAvatarUseCase{
		UserRepository:   userRepository,
		UploadRepository: uploadRepository,
	}
}

func (uc *UserAvatarUseCase) Execute(ctx context.Context, dto UserAvatarDTO) error {
	user, err := uc.UserRepository.FindByID(ctx, dto.UserID)
	if err != nil {
		return err
	}

	if user == nil {
		return errorformatted.BadRequestError(trace.GetTrace(), "user_not_found", "User not found")
	}

	avatarURL, err := uc.UploadRepository.Upload(ctx, dto.Avatar, user.ID)
	if err != nil {
		return err
	}
	user.DefineAvatar(avatarURL)

	err = uc.UserRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil

}
