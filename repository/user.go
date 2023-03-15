package repository

import (
	"context"
	"gorm.io/gorm"
	"todo-web/entity"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
	GetUserByEmailNUsername(ctx context.Context, username string, email string) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) error
	AddUser(ctx context.Context, user entity.User) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).First(&user).Where("username = ?", username).Error; err != nil {
		return entity.User{}, err
	}
	return entity.User{}, nil
}

func (r *userRepository) GetUserByEmailNUsername(ctx context.Context, username string, email string) (entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int) (entity.User, error) {

	return entity.User{}, nil
}

func (r *userRepository) AddUser(ctx context.Context, user entity.User) (entity.User, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) error {
	return nil
}
