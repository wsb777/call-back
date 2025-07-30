package services

import (
	"context"
	"errors"

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
	userId := claims.UserID

	revoked, err := s.jwtRepo.IsTokenRevoked(ctx, claims.TokenID)
	if err != nil {
		return "", "", err
	}
	if revoked {
		return "", "", errors.New("token revoked")
	}
	accessToken, refreshToken, err := s.jwtEncoder.GenerateTokenPair(userId)

	if err != nil {
		return "", "", err
	}

	err = s.jwtRepo.RevokeToken(ctx, &models.JwtToken{
		ID:        claims.TokenID,
		UserId:    userId,
		ExpiresAt: claims.ExpiresAt.Time,
	})

	if err != nil {
		return "", "", errors.New("failed to revoke old token")
	}

	if err != nil {
		return "", "", err
	}

	go func() {
		_ = s.jwtRepo.CleanupExpiredTokens(context.Background())
	}()

	return accessToken, refreshToken, nil
}
