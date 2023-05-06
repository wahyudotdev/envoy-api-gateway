package helper

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

const jwtSecret = "fe9767d5dd989526d4648b0e3d22cc177e1356bf"

func SignJwt(id string, roles []string) (string, error) {

	duration := time.Minute * 60
	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"exp":   time.Now().Add(duration).Unix(),
		"id":    id,
		"roles": strings.Join(roles, ","),
	})

	token, err := sign.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseJwt(token string) (jwt.MapClaims, error) {
	data, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid token")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := data.Claims.(jwt.MapClaims)
	if !ok || !data.Valid {
		return nil, errors.New("invalid token")
	}

	exp := claims["exp"].(float64)
	if time.Now().Unix() > int64(exp) {
		return nil, errors.New("expired token")
	}
	return claims, nil
}
