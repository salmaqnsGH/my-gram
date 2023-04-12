package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = []byte("SECRET-KEY")

func GenerateToken(userID uint, email string) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["email"] = email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errors.New("Bearer token not found")
	}

	stringToken := headerToken[7:]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Failed to get sign token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, errors.New("Invalid token")
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, errors.New("Failed to parse claims")
	}

	return token.Claims.(jwt.MapClaims), nil
}
