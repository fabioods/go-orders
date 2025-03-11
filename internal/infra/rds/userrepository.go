package rds

import (
	"context"

	"github.com/fabioods/go-orders/internal/entity"
)

type UserRepository struct{}

func NewUserRepositoryRDS() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(ctx context.Context, user *entity.User) error {
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	return nil
}
