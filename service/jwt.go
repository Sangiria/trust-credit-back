package service

import (
	"errors"
	"time"
	"trust-credit-back/environment"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

//надо ли отправлять возможные ошибки, произошедшие во время подписи токена, и регулировать их в ручках где происходит генерация jwt токенов?

func NewToken(id string, is_access bool) string {
	var (
		token_secret string
		duration time.Duration
	)

	if is_access {
		token_secret = environment.GetVariable("ACCESS_SECRET")
		duration = time.Minute * 15
	} else {
		token_secret = environment.GetVariable("REFRESH_SECRET")
		duration = time.Hour * 168
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(duration).Unix(),
	})

	signed_token, err :=  token.SignedString([]byte(token_secret))
	
	if err != nil {
		return ""
	}

	return signed_token
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