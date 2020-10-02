package router

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kangbojk/go-react-fullstack/pkg/ID"
	"time"
)

var mySigningKey = []byte("secret-xd")

type MyCustomClaims struct {
	UserEmail string `json:"username"`
	jwt.StandardClaims
}

func gen_token(userEmail string, userID id.ID) (string, error) {

	/* Set token claims */

	// Create the Claims
	claims := MyCustomClaims{
		userEmail,
		jwt.StandardClaims{
			Subject:   id.IDToString(userID),
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	/* Sign the token with our secret */
	return token.SignedString(mySigningKey)
}

//ValidateToken will validate the token
func ValidateToken(myToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(myToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if err != nil {
		return false, ""
	}

	claims := token.Claims.(*MyCustomClaims)
	return token.Valid, claims.UserEmail
}

// func ParseToken(myToken string) (*Token, error) {

// }
