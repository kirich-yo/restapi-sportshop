package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(username string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
	})

	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return s, nil
}
