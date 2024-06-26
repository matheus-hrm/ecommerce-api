package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/matheus-hrm/trampos/config"
)
func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpiration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),	
	})

	tokenstring, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenstring, err
}