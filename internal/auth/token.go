package auth

import (
	"os"

	"github.com/golang-jwt/jwt"
)

type MasterClaim struct {
	Address string `json:"address"` //ip:port
	jwt.StandardClaims
}

func NewAccessToken(claims MasterClaim) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ParseAccessToken(accessToken string) *MasterClaim {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &MasterClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return parsedAccessToken.Claims.(*MasterClaim)
}
