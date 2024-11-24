package parsejwt

import (
	"context"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

const secretKey = "DMS_Microservices_system1"

type TokenClaims struct {
	jwt.StandardClaims
	UserID int64
}

func ParseToken(ctx context.Context, accessToken string) (int64, error) {

	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserID, nil
}
