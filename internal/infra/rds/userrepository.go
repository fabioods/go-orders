package rds

import (
	"context"

	"github.com/fabioods/go-orders/internal/entity"
	"github.com/fabioods/go-orders/internal/errorcode"
	"github.com/fabioods/go-orders/pkg/errorformatted"
	"github.com/fabioods/go-orders/pkg/trace"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepositoryRDS(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&entity.User{})
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, user *entity.User) error {
	userORM := r.db.Create(user)
	if userORM.Error != nil {
		err := errorformatted.UnexpectedError(trace.GetTrace(), errorcode.ErrorSaveUserError, "Failed to save user: %v", userORM.Error)
		return err
	}

	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		err := errorformatted.UnexpectedError(trace.GetTrace(), errorcode.ErrorFindByIdError, "Failed to find user by id: %v", result.Error)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("email = ? ", email).First(&user)
	if result.Error != nil {
		err := errorformatted.UnexpectedError(trace.GetTrace(), errorcode.ErrorFindByEmailError, "Failed to find user by email: %v", result.Error)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	err := r.db.Save(user).Error
	if err != nil {
		errFmt := errorformatted.UnexpectedError(trace.GetTrace(), errorcode.ErrorUpdateUserError, "Failed to update user: %v", err)
		return errFmt
	}
	return nil
}
