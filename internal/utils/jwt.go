package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.RegisteredClaims
}

var (
	jwtSecret       []byte
	accessTokenExp  time.Duration
	refreshTokenExp time.Duration
)

func SetJWTConfig(secret string, accessExp, refreshExp time.Duration) {
	jwtSecret = []byte(secret)
	accessTokenExp = accessExp
	refreshTokenExp = refreshExp
}

func GenerateAccessToken(userID uint, username string, firstName string, lastName string) (string, error) {
	if len(jwtSecret) == 0 {
		return "", errors.New("JWT secret is not set")
	}

	claims := JWTClaims{
		UserID:    userID,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(userID uint, username string, firstName string, lastName string) (string, error) {
	if len(jwtSecret) == 0 {
		return "", errors.New("JWT secret is not set")
	}

	claims := JWTClaims{
		UserID:    userID,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*JWTClaims, error) {
	if len(jwtSecret) == 0 {
		return nil, errors.New("JWT secret is not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
