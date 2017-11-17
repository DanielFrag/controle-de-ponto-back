package utils

import (
	"time"
	"fmt"
	"os"
	"strconv"
	jwt "github.com/dgrijalva/jwt-go"
)

func TokenChecker(token string) (interface{}, error) {
	tokenParsed, parseError := jwt.Parse(token, func(tk *jwt.Token) (interface {}, error) {
		if _, success := tk.Method.(*jwt.SigningMethodHMAC); !success {
			return nil, fmt.Errorf("Error. Invalid signed token")
		}
		tkSecret := os.Getenv("TOKEN_SECRET")
		if tkSecret == "" {
			tkSecret = "secret"
		}
		return []byte(tkSecret), nil
	})
	if parseError != nil {
		return nil, parseError
	}
	if claims, ok := tokenParsed.Claims.(jwt.MapClaims); ok {
		return claims["data"], parseError
	}
	return nil, parseError
}

func EncodeToken(payload interface{}) (string, error) {
	tkExpTime := os.Getenv("TOKEN_EXPIRATION_TIME")
	if tkExpTime == "" {
		tkExpTime = "60"
	}
	seconds, strConvError := strconv.ParseInt(tkExpTime, 10, 64)
	if strConvError != nil {
		seconds = 60
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"data": payload,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * time.Duration(seconds)).Unix(),
	})
	tkSecret := os.Getenv("TOKEN_SECRET")
	if tkSecret == "" {
		tkSecret = "secret"
	}
	return token.SignedString([]byte(tkSecret))
}