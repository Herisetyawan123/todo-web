package service

import (
	"context"
	"todo-web/entity"
	"todo-web/repository"
)

type UserService interface {
	GetUser(ctx context.Context, id int) (entity.User, error)
	Login(ctx context.Context, data entity.LoginRequest) error
	Register(ctx context.Context, data entity.User) error
	Token(ctx context.Context, id int) (string, error)
	Verification(ctx context.Context, id int) error
}
type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository}
}

func (s *userService) GetUser(ctx context.Context, id int) (entity.User, error) {
	return entity.User{}, nil
}

func (s *userService) Login(ctx context.Context, data entity.LoginRequest) error {
	// TODO: service login
	_, err := s.userRepository.GetUserByUsername(ctx, data.Username)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) Register(ctx context.Context, data entity.User) error {
	// TODO: service login
	err := s.userRepository.AddUser(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) Token(ctx context.Context, id int) (string, error) {
	// TODO: get token from this service
	return "", nil
}

func (s *userService) Verification(ctx context.Context, id int) error {
	return nil
}
