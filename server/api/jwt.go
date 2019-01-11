package api

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// Claims ...
type Claims struct {
	Username string
	Provider string
	jwt.StandardClaims
}

// JWTHandler ...
type JWTHandler interface {
	GetToken(Claims) (string, error)
	GetClaims(tokenString string) (*Claims, error)
}

// NewJWTHandler ...
func NewJWTHandler(pkey *rsa.PrivateKey) JWTHandler {
	return &jwtHandler{
		pKey: pkey,
	}
}

type jwtHandler struct {
	secret string
	pKey   *rsa.PrivateKey
}

func (h *jwtHandler) GetToken(claims Claims) (string, error) {
	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	tokenString, e := t.SignedString(h.pKey)
	return tokenString, e
}

func (h *jwtHandler) GetClaims(tokenString string) (*Claims, error) {
	t, e := jwt.ParseWithClaims(tokenString, &Claims{},
		func(x *jwt.Token) (a interface{}, d error) {
			a, d = &h.pKey.PublicKey, nil
			return
		},
	)
	if e != nil {
		return nil, e
	}
	claims, ok := t.Claims.(*Claims)
	if !ok {
		panic(
			fmt.Errorf("False JWTUser type assertion. Security breach. Private key compromised"),
		)
	}
	return claims, nil
}
