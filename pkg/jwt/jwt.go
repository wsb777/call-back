package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/wsb777/call-back/internal/config"
)

type Encoder interface {
	CreateAccessToken(userId string) (string, error)
	VerifyToken(tokenString string) (*AccessClaims, error)
	VerifyRefreshToken(tokenString string) (*RefreshClaims, error)
	GenerateTokenPair(userId string) (accessToken, refreshToken string, err error)
}

type AccessClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserID  string `json:"user_id"`
	TokenID string `json:"token_id"`
	jwt.RegisteredClaims
}

type JWTEncoder struct {
	secret string
}

func NewJWTEncoder(cfg *config.Config) *JWTEncoder {
	return &JWTEncoder{secret: cfg.JWTSecret}
}

func (j *JWTEncoder) CreateAccessToken(userId string) (string, error) {
	key := []byte(j.secret)
	claims := AccessClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // Срок действия
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "call-back",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}

func (j *JWTEncoder) CreateRefreshToken(userId string) (string, error) {
	key := []byte(j.secret)
	tokenID := uuid.New().String()
	claims := RefreshClaims{
		userId,
		tokenID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // Срок действия
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "call-back",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}

func (j *JWTEncoder) GenerateTokenPair(userId string) (accessToken, refreshToken string, err error) {
	accessToken, err = j.CreateAccessToken(userId)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = j.CreateRefreshToken(userId)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (j *JWTEncoder) VerifyToken(tokenString string) (*AccessClaims, error) {
	key := []byte(j.secret)
	token, err := jwt.ParseWithClaims(
		tokenString,
		&AccessClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		},
	)

	if claims, ok := token.Claims.(*AccessClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func (j *JWTEncoder) VerifyRefreshToken(tokenString string) (*RefreshClaims, error) {
	key := []byte(j.secret)
	token, err := jwt.ParseWithClaims(
		tokenString,
		&RefreshClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		},
	)

	if claims, ok := token.Claims.(*RefreshClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func TestJWTVerification() {
	encoder := &JWTEncoder{secret: "your_secret_key"}

	// Создание тестового токена
	token, _ := encoder.CreateAccessToken("test_user")
	log.Printf("Generated token: %s", token)

	// Проверка токена
	claims, err := encoder.VerifyToken(token)
	if err != nil {
		log.Printf("Verification error: %v", err)
	} else {
		log.Printf("Verified claims: %+v", claims)
	}
}
