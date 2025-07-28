package services

import (
	"errors"
	"fmt"

	"github.com/wsb777/call-back/internal/db/repo"
	"github.com/wsb777/call-back/internal/dto"
	"github.com/wsb777/call-back/pkg/hasher"
	_jwt "github.com/wsb777/call-back/pkg/jwt"
)

type AuthService interface {
	SignIn(req dto.AuthDto) (string, string, error)
}

type authService struct {
	userRepo   repo.UserRepo
	hasher     hasher.PasswordHasher
	jwtEncoder _jwt.Encoder
}

func NewAuthService(repo repo.UserRepo, hasher hasher.PasswordHasher, jwtEncoder _jwt.Encoder) AuthService {
	return &authService{userRepo: repo, hasher: hasher, jwtEncoder: jwtEncoder}
}

func (s *authService) SignIn(req dto.AuthDto) (string, string, error) {
	if existing, _ := s.userRepo.FindByLogin(req.Password); existing != nil {
		if s.hasher.ComparePassword(existing.Password, req.Password) {
			id := fmt.Sprintf("%d", existing.ID)
			accessToken, refreshToken, err := s.jwtEncoder.GenerateTokenPair(id)
			if err != nil {
				return "", "", errors.New("Проблема с созданием JWT")
			}
			return accessToken, refreshToken, nil
		}
		return "", "", errors.New("Пароль не подошел")
	}
	return "", "", errors.New("Пользователя не существует")
}
