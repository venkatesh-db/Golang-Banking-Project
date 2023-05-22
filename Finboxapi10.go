package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {

	fmt.Println("Hello World!")

	mySigningKey := []byte("coderrange")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "venkatesh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(mySigningKey)

	fmt.Printf("%v %v", ss, err)
}