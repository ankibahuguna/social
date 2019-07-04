package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWT(userid uint, hours time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["uid"] = userid
	claims["exp"] = time.Now().Add(time.Hour * hours).Unix()

	jwtToken, err := token.SignedString([]byte("thisIsASafeSe2374823478#$$%$%^key"))

	return jwtToken, err
}
