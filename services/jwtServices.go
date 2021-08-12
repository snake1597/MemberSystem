package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	SecretKey = []byte("arthur")
)

func GenerateToken(id int, account string) (authToken string, err error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["account"] = account
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	authToken, err = token.SignedString(SecretKey)
	return
}

func ValidateToken(token string) (c map[string]interface{}, err error) {

	authToken, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("token not valild")
	}

	if claim, ok := authToken.Claims.(jwt.MapClaims); ok && authToken.Valid {
		c = map[string]interface{}{
			"id":      claim["id"],
			"account": claim["account"],
		}
		return c, nil
	}

	return nil, fmt.Errorf("token not valild")
}
