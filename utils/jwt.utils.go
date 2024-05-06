package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	Id uint64
}

func ProvideJWT(payload *Payload) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
		"userId": payload.Id,
	}
	claims["userId"] = strconv.FormatUint(claims["userId"].(uint64), 10)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}
