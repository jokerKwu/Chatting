package model

import "github.com/dgrijalva/jwt-go"

//jwt session data
type TokenData struct{
	Name string `json:"name" validate:"required"`
}


type TokenMetaData struct{
	IssuedAt int64
	ExpiredAt int64
}

//claim 데이타
type JwtClaim struct{
	Name string `json:"name" validate:"required"`
	jwt.StandardClaims
}

