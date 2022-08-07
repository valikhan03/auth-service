package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTWrapper struct {
	SecretKey      string
	Issuer         string
	ExpirationTime time.Duration
}

type JWTClaims struct {
	Email string
	jwt.StandardClaims
}

func (j *JWTWrapper) GenerateToken(email string) (signedToken string, err error) {
	claims := JWTClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Issuer: j.Issuer,
			ExpiresAt: time.Now().Add(j.ExpirationTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	return
}


func (j *JWTWrapper) ValidateToken(signedToken string) (claims *JWTClaims, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", t.Header["alg"]))
		}

		return []byte(j.SecretKey), nil
	})

	if err != nil{
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); token.Valid && ok{
		if claims.ExpiresAt < time.Now().Local().Unix(){
			err = errors.New("JWT is expired")
		}
	}
	
	return
}