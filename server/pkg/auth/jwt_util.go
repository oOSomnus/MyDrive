package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = "invalid token"
)

type JWTManager struct {
	secretKey string
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{secretKey: secretKey}
}

func (jm *JWTManager) GenerateToken(userID uint, expiration time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiration)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jm.secretKey))
}

func (jm *JWTManager) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jm.secretKey), nil
		},
	)

	if err != nil {
		return nil, errors.New(ErrInvalidToken)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New(ErrInvalidToken)
}
