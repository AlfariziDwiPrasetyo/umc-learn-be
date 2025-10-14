package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func CreateToken(id int64, email string, secretKey string) (*Tokens, error) {
	key := []byte(secretKey)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(10 * time.Minute).Unix(),
	})
	accessStr, err := accessToken.SignedString(key)
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	refreshStr, err := refreshToken.SignedString(key)

	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:  accessStr,
		RefreshToken: refreshStr,
	}, nil
}

func ValidateToken(tokenStr string, secretKey string) (int64, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	return int64(claims["id"].(float64)), nil
}
