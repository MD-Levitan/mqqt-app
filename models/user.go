package models

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserClaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type Error struct {
	Error string `json:"error"`
}

type JWT struct {
	JWT string `json:"Bearer"`
}

func CreateUserFromJWT(jwtClaim UserClaim) *User {
	u := &User{jwtClaim.Username, jwtClaim.Password}
	return u
}
