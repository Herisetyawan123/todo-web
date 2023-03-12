package service

import (
	"context"
	"errors"
	"todo-web/entity"
	"todo-web/repository"
)

type UserService interface {
	GetUser(ctx context.Context, id int) (entity.User, error)
	Login(ctx context.Context, data entity.LoginRequest) error
	Register(ctx context.Context, data *entity.User) (entity.User, error)
	Token(ctx context.Context, id int) (string, error)
	Verification(ctx context.Context, id int) error
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository}
}

func (s *userService) GetUser(ctx context.Context, id int) (entity.User, error) {
	return entity.User{}, nil
}

func (s *userService) Login(ctx context.Context, data entity.LoginRequest) error {
	// TODO: service login
	_, err := s.repo.GetUserByUsername(ctx, data.Username)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) Register(ctx context.Context, data *entity.User) (entity.User, error) {
	_, err := s.repo.GetUserByEmailNUsername(ctx, data.Username, data.Password)
	if err == nil {
		return entity.User{}, errors.New("Username atau email sudah ada coba pilih username atau email yg berbeda")
	}

	newUser, err := s.repo.AddUser(ctx, *data)
	if err != nil {
		return entity.User{}, err
	}
	return newUser, nil
}

func (s *userService) Token(ctx context.Context, id int) (string, error) {
	// TODO: get token from this service
	return "", nil
}

func (s *userService) Verification(ctx context.Context, id int) error {
	return nil
}
