package service

import (
	"errors"
	"time"
	"trust-credit-back/environment"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

var (
	access_secret = environment.GetVariable("ACCESS_SECRET")
	access_duration = time.Minute * 15
	refresh_secret = environment.GetVariable("REFRESH_SECRET")
	refresh_duration = time.Hour * 168
)

func NewTokens(id string) (TokenPair, error) {
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(access_duration).Unix(),
	})

	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(refresh_duration).Unix(),
	})

	signed_access, err1 :=  access_token.SignedString([]byte(access_secret))
	signed_refresh, err2 :=  refresh_token.SignedString([]byte(refresh_secret))
	
	if err1 != nil || err2 != nil {
		return TokenPair{}, errors.Join(err1, err2)
	}

	return TokenPair{AccessToken: signed_access, RefreshToken: signed_refresh}, nil
}

func ParseToken(token string, secret string) (*jwt.Token, error) {
	parsed_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return parsed_token, err
}

func ValidateToken(token_string string, secret string) (string, error) {
	token, err := ParseToken(token_string, secret)
	if err != nil || token == nil || !token.Valid {
		return "", errors.New("token invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	id, ok := claims["id"].(string)
	_, err = uuid.Parse(id)
	
	if !ok {
		return "", errors.New("invalid token claims")
	}

	if err != nil {
		return "", errors.New("invalid format")
	}
	
	exp, ok := claims["exp"].(float64)
	if ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return "", errors.New("token expired")
		}
	} else {
		return "", errors.New("invalid or missing expiration time")
	}

	return id, nil
}