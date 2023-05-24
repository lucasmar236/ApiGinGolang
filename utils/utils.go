package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
	"time"
)

var secretkey = []byte("opa")

func GenerateJwt(user string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = user
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	tokenString, err := token.SignedString(secretkey)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return tokenString, nil
}

func VerifyJwt(token string) error {
	token = strings.ReplaceAll(token, "Bearer ", "")
	tokenParse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("erro ao validar o token")
		}
		return secretkey, nil
	})
	if err != nil {
		return err
	}
	if !tokenParse.Valid {
		fmt.Println(tokenParse.Valid)
		return errors.New("token error")
	} else {
		return nil
	}
}
