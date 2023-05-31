// github.com/bartmika/mulberry-server/pkg/utils/jwt.go
package utils

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// Generate the `access token` and `refresh token` for the secret key.
func GenerateJWTTokenPair(secretKey []byte, sessionUUID string) (string, string, error) {
	//
	// Generate token.
	//
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["session_uuid"] = sessionUUID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	//
	// Generate refresh token.
	//
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["session_uuid"] = sessionUUID
	rtClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}

// Validates either the `access token` or `refresh token` and returns either the
// `user_uuid` if success or error on failure.
func ProcessJWTToken(hmacSecret []byte, reqToken string) (string, error) {
	token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err == nil && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			session_uuid := claims["session_uuid"].(string)
			return session_uuid, nil
		} else {
			return "", err
		}

	} else {
		return "", err
	}
}
