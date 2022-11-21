package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	secret      = "secretkeyuntukfinalproject"
	jwtDuration = 1 * time.Hour
)

type CustomClaims struct {
	UserId string `json:"user_id"`
	RoleId int    `json:"role_id"`
	jwt.RegisteredClaims
}

func GenerateToken(user_id string, role_id int) (string, error) {
	now := time.Now()

	claims := CustomClaims{
		user_id,
		role_id,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(jwtDuration)),
			Issuer:    "jwtissuer",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ValidateToken(input string) (string, int, error) {
	token, err := jwt.ParseWithClaims(
		input,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if err != nil {
		return "", -1, nil
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.UserId, claims.RoleId, nil
	} else {
		return "", -1, err
	}
}
