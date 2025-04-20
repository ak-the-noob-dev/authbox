package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type CustomClaims struct {
	UserID uint      `json:"uid"`
	Type   TokenType `json:"type"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

// CreateToken creates a JWT token with a given TTL (in minutes)
func CreateToken(secret string, userID uint, role string, ttlMinutes int, tokenType TokenType) (string, error) {
	now := time.Now()
	claims := CustomClaims{
		UserID: userID,
		Type:   tokenType,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(ttlMinutes) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken verifies and extracts claims
func ParseToken(secret, tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
