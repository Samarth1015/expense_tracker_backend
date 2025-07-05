package service

import (
	"fmt"
	"log"
	"os"
	"time"

	jwttoken "github.com/Samarth1015/expense/dto/jwtToken"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func CreateToken(data jwttoken.JwtToken) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       data.Id,
			"username": data.Username,
			"role":     data.Role,
			"email":    data.Email,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func VerifyJWTToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

var c jwttoken.Claims

func ClaimToken(tokenString string) (jwttoken.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		log.Println("error:", "unable to parse the token", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if email, ok := claims["email"].(string); ok {
			c.Email = email
		}
		if id, ok := claims["id"].(string); ok {
			c.ID = id
		}
		if role, ok := claims["role"].(string); ok {
			c.Role = role
		}
		if username, ok := claims["username"].(string); ok {
			c.UserName = username
		}
		return c, nil
	}

	return jwttoken.Claims{}, err
}
