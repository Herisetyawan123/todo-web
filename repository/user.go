package repository

import (
	"context"
	"gorm.io/gorm"
	"todo-web/entity"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int) (entity.User, error)
	GetUserByUsername(ctx context.Context, email string) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) error
	AddUser(ctx context.Context, user entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByUsername(ctx context.Context, email string) (entity.User, error) {
	return entity.User{}, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int) (entity.User, error) {
	return entity.User{}, nil
}

func (r *userRepository) AddUser(ctx context.Context, user entity.User) error {
	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) error {
	return nil
}
