package utils

import (
	"errors"
	"fmt"
	"log"
	"muxWithSql/config"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var User interface{}

func CreateToken(claims *config.JWTToken) (string, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("ENV not loaded")
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims.RegisteredClaims).SignedString([]byte(os.Getenv("SECRET_JWT")))

	if err != nil {
		fmt.Println("JWT Creation failed")
		return "", err
	}

	return token, nil

}

func ValidateToken(tokens string) (interface{}, error) {
	
	token, _ := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		if _ , ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_JWT")), nil
	})

	claim, ok := token.Claims.(jwt.MapClaims)
	
	fmt.Println("user Token comming in",token.Valid)
	if ok && token.Valid {
		User = claim["sub"]
	} else if float64(time.Now().Unix()) > claim["exp"].(float64) && token.Valid {
		return nil, errors.New("JWT expiry")
	} else {
		return nil, errors.New("Invalid token")
	}
	
	return User, nil

}