package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (string, error) {
	
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	fmt.Println(config.JwtSecret)
	return token.SignedString([]byte(config.JwtSecret))
}

func ValidateToken(r *http.Request) error {

	tokenString := extractToken(r)

	token, error := jwt.Parse(tokenString, returnJwtSecret)

	if  error != nil {
		return error
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func ExtractUserId(r *http.Request) (uint64, error) {

	tokenString := extractToken(r)

	token, error := jwt.Parse(tokenString, returnJwtSecret)

	if  error != nil {
		return 0, error
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, error := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)

		if  error != nil {
			return 0, error
		}

		return userId, nil
	}

	return 0, errors.New("token inválido")

}

func returnJwtSecret(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected JWT Signin Method %v", token.Header["alg"])
	}

	return []byte(config.JwtSecret), nil
}
