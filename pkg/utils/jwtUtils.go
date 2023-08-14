package utils

import (
	"BookHive/pkg/types"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userData types.UserData) string {
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	secretKey := []byte(config.AccessTokenSecret)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = userData.Id
	claims["username"] = userData.Username
	claims["admin"] = userData.Admin
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}

	return tokenStr
}

func DecodeJWT(tokenStr string) (jwt.MapClaims, error) {
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	secretKey := []byte(config.AccessTokenSecret)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
