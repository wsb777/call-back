package services

import (
	"errors"
	"fmt"

	"github.com/wsb777/call-back/internal/db/repo"
	"github.com/wsb777/call-back/internal/dto"
	"github.com/wsb777/call-back/pkg/hasher"
)

type UserSignInService interface {
	SignIn(req dto.UserSignInDto) (string, error)
}

type userSignInService struct {
    userRepo repo.UserRepo
    hasher   hasher.PasswordHasher
}

func NewUserSignInService(repo repo.UserRepo, hasher hasher.PasswordHasher) UserSignInService {
	return &userSignInService{userRepo: repo, hasher: hasher}
}

func (s *userSignInService) SignIn(req dto.UserSignInDto) (string, error) {
	if existing, _ :=s.userRepo.FindByLogin(req.Password); existing != nil {
		if s.hasher.ComparePassword(existing.Password, req.Password) {
			return fmt.Sprintf("Идентификатор: %v", existing.ID), nil
		}
		return "", errors.New("Пароль не подошел")
	}
	return "", errors.New("Пользователя не существует")
}

