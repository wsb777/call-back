package services

import (
	"context"
	"errors"
	"log"

	"github.com/wsb777/call-back/internal/db/repo"
	"github.com/wsb777/call-back/internal/dto"
	"github.com/wsb777/call-back/internal/models"
	"github.com/wsb777/call-back/pkg/hasher"
	_jwt "github.com/wsb777/call-back/pkg/jwt"
)

type RefreshService interface {
	RefreshToken(ctx context.Context, token dto.RefreshTokenDto) (string, string, error)
}

type refreshService struct {
	jwtRepo    repo.JWTRepo
	hasher     hasher.PasswordHasher
	jwtEncoder _jwt.Encoder
}

func NewRefreshService(jwtRepo repo.JWTRepo, hasher hasher.PasswordHasher, jwtEncoder _jwt.Encoder) RefreshService {
	return &refreshService{jwtRepo: jwtRepo, hasher: hasher, jwtEncoder: jwtEncoder}
}

func (s *refreshService) RefreshToken(ctx context.Context, token dto.RefreshTokenDto) (string, string, error) {
	claims, err := s.jwtEncoder.VerifyRefreshToken(token.Token)
	if err != nil {
		return "", "", err
	}
	if claims == nil {
		return "", "", errors.New("В claims ничего нету")
	}
	log.Println(claims)
	userId := claims.UserID
	accessToken, refreshToken, err := s.jwtEncoder.GenerateTokenPair(userId)
	if err != nil {
		return "", "", err
	}
	modelRefreshToken, err := s.jwtEncoder.VerifyRefreshToken(refreshToken)

	if err != nil {
		return "", "", err
	}

	err = s.jwtRepo.CleanupExpiredTokens(ctx)

	if err != nil {
		return "", "", err
	}

	err = s.jwtRepo.RevokeToken(ctx, &models.JwtToken{ID: modelRefreshToken.TokenID, UserId: modelRefreshToken.UserID, ExpiresAt: modelRefreshToken.ExpiresAt.Time})

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
