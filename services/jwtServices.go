package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	jwtKey string
)

func init() {

	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	jwtKey = viper.GetString("jwtKey")
}

func GenerateToken(id int, account string) (authToken string, err error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["account"] = account
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	authToken, err = token.SignedString([]byte(jwtKey))
	return
}

func ValidateToken(token string) (c map[string]interface{}, err error) {

	authToken, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse token error")
	}

	if claim, ok := authToken.Claims.(jwt.MapClaims); ok {
		c = map[string]interface{}{
			"id":      claim["id"],
			"account": claim["account"],
		}
		return c, nil
	}

	return nil, fmt.Errorf("token is not valild")
}
