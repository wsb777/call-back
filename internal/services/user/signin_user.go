package services

import (
	"errors"
	"fmt"

	"github.com/wsb777/call-back/internal/db/repo"
	"github.com/wsb777/call-back/internal/dto"
	"github.com/wsb777/call-back/pkg/hasher"
	_jwt "github.com/wsb777/call-back/pkg/jwt"
)

type UserSignInService interface {
	SignIn(req dto.UserSignInDto) (string, error)
}

type userSignInService struct {
	userRepo   repo.UserRepo
	hasher     hasher.PasswordHasher
	jwtEncoder _jwt.Encoder
}

func NewUserSignInService(repo repo.UserRepo, hasher hasher.PasswordHasher, jwtEncoder _jwt.Encoder) UserSignInService {
	return &userSignInService{userRepo: repo, hasher: hasher, jwtEncoder: jwtEncoder}
}

func (s *userSignInService) SignIn(req dto.UserSignInDto) (string, error) {
	if existing, _ := s.userRepo.FindByLogin(req.Password); existing != nil {
		if s.hasher.ComparePassword(existing.Password, req.Password) {
			id := fmt.Sprintf("%d", existing.ID)
			token, err := s.jwtEncoder.CreateToken(id)
			if err != nil {
				return "", errors.New("Проблема с созданием JWT")
			}
			return token, nil
		}
		return "", errors.New("Пароль не подошел")
	}
	return "", errors.New("Пользователя не существует")
}
