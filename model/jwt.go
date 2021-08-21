package model

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	ID        string `json:"id" xml:"id"`
	Email     string `json:"email" xml:"email"`
	Name      string `json:"name" xml:"name"`
	Lastname  string `json:"lastname" xml:"lastname"`
	BirthDate string `json:"birthDate" xml:"birthDate"`
	Biography string `json:"biography" xml:"biography"`
	Location  string `json:"location" xml:"location"`
	WebSite   string `json:"webSite" xml:"webSite"`
	jwt.StandardClaims
}

type Token struct {
	Token string `json:"access_token" xml:"access_token"`
}
