package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	helper "github.com/dgrijalva/jwt-go/test"
)

func main() {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().UTC().UnixNano(),
		"exp": time.Now().UTC().Add(time.Second * 60).UnixNano(),
	})

	privateKey := helper.LoadRSAPrivateKeyFromDisk("./jwtRSA256-private.pem")

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(privateKey)
	fmt.Printf("%s\n%s\n", tokenString, err)
}
