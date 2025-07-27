package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wsb777/call-back/internal/config"
)

type Encoder interface {
	CreateToken(userId string) (string, error)
	VerifyToken(tokenString string) (*Claims, error)
}

type Claims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

type JWTEncoder struct {
	secret string
}

func NewJWTEncoder(cfg *config.Config) *JWTEncoder {
	log.Printf("NewJWTEncoder CALLED with secret: %s", cfg.JWTSecret)
	return &JWTEncoder{secret: cfg.JWTSecret}
}

func (j *JWTEncoder) CreateToken(userId string) (string, error) {
	key := []byte(j.secret)
	log.Printf("CreateToken: j=%p, secret='%s'", j, j.secret)
	claims := Claims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Срок действия
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "call-back",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}

func (j *JWTEncoder) VerifyToken(tokenString string) (*Claims, error) {
	key := []byte(j.secret)
	log.Printf("VerifyToken: using secret: %s", j.secret)
	log.Println(tokenString)
	log.Println("Проверка токена...")
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		},
	)

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func TestJWTVerification() {
	encoder := &JWTEncoder{secret: "your_secret_key"}

	// Создание тестового токена
	token, _ := encoder.CreateToken("test_user")
	log.Printf("Generated token: %s", token)

	// Проверка токена
	claims, err := encoder.VerifyToken(token)
	if err != nil {
		log.Printf("Verification error: %v", err)
	} else {
		log.Printf("Verified claims: %+v", claims)
	}
}
