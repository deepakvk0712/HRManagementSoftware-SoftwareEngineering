package models

import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	Email         string
	StandardClaim jwt.StandardClaims
	Role          byte
}

type JWTParams struct {
	SecretKey    string
	Issuer       string
	LoginTimeout int
}

// Implemented this to avoid error while generating token with claim
func (J JWTClaim) Valid() error {
	//TODO implement me
	//panic("implement me")
	return nil
}
