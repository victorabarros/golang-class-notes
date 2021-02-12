package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	helper "github.com/dgrijalva/jwt-go/test"
)

var (
	expiration = flag.Int("exp", 60, "token expiration (seconds)")
	keyPath    = flag.String("key", "./jwtRSA256-private.pem", "path for the private key .pem")
	headerKid  = flag.String("kid", "kid", "header kid")
)

func main() {
	flag.Parse()
	tokenString, err := buildToken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tokenString)
}

func buildToken() (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().UTC().Unix(),
		"exp": time.Now().UTC().Add(time.Second * time.Duration(*expiration)).Unix(),
	})
	token.Header["kid"] = *headerKid

	privateKey := helper.LoadRSAPrivateKeyFromDisk(*keyPath)

	// Sign and get the complete encoded token as a string
	return token.SignedString(privateKey)
}
