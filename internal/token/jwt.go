package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const tokenKey = "sisi_ieches"

type TokenClaims struct {
	Email string `json:"email"`
	Exp   int64  `json:"exp"`
}

func (t TokenClaims) Valid() error {
	if time.Now().Unix() > t.Exp {
		return errors.New("token has expired")
	}
	return nil
}

func New(claims TokenClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": claims.Email,
		"exp":   time.Now().Add(time.Millisecond * time.Duration(claims.Exp)).Unix(),
	}).SignedString([]byte(tokenKey))
}

func ExtractClaims(tokenStr string) (*TokenClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenKey), nil
	}

	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
