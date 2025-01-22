package service

import (
    "context"
    "fmt"

    "github.com/google/uuid"
    "github.com/nurdamiron/printer-automation/internal/domain"
    "github.com/nurdamiron/printer-automation/internal/repository"
)

type UserService interface {
    CreateUser(ctx context.Context, username string) (*domain.User, error)
    GetUser(ctx context.Context, id string) (*domain.User, error)
}

type userService struct {
    userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{userRepo: repo}
}

func (s *userService) CreateUser(ctx context.Context, username string) (*domain.User, error) {
    user := &domain.User{
        ID:       uuid.New(),
        Username: username,
    }
    err := s.userRepo.Create(ctx, user)
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }
    return user, nil
}

func (s *userService) GetUser(ctx context.Context, id string) (*domain.User, error) {
    return s.userRepo.GetByID(ctx, id)
}
