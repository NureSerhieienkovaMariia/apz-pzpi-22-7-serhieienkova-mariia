package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour

	hostEmail   = "25sergeenkova.maria@gmail.com"
	appPassword = "yzvoglpnmpdqanuq"

	smtpHost = "smtp.gmail.com"
	smtpPort = "587"
)

type TokenClaims struct {
	jwt.StandardClaims
	Id       int    `json:"id"`
	UserType string `json:"user_type"`
}

func ParseToken(accessToken string) (TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return TokenClaims{}, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return TokenClaims{}, errors.New("token claims are not of type *TokenClaims")
	}

	return *claims, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
