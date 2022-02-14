package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 秘钥
var jwtSecret []byte

// TokenExpireDuration 过期时间
const TokenExpireDuration = time.Hour * 2 //设置过期时间

type Claims struct {
	UsesId int `json:"userId"`
	jwt.StandardClaims
}

func GenerateToken(userId int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(TokenExpireDuration)

	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "designer",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func RefreshToken(token string) (string, error) {
	claims, err := ParseToken(token)
	if err != nil {
		return "", err
	}

	return GenerateToken(claims.UsesId)
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
