package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	helper "github.com/dgrijalva/jwt-go/test"
)

func main() {
	tokenString, err := buildToken()
	fmt.Printf("%s\n%s\n", tokenString, err)
}

func buildToken() (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().UTC().Unix(),
		"exp": time.Now().UTC().Add(time.Second * 60).Unix(),
	})

	privateKey := helper.LoadRSAPrivateKeyFromDisk("./jwtRSA256-private.pem")

	// Sign and get the complete encoded token as a string
	return token.SignedString(privateKey)
}
