package tool

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type TypeClaims struct {
	UserId int
	jwt.StandardClaims
}

var hmacSampleSecret []byte

func CreateToken() (string, int64) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":     1,
		"expireTime": expireTime,
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return tokenString, expireTime.Unix()
}

func ParseToken(tokenString string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println(claims["userId"], claims["expireTime"])
	} else {
		fmt.Println(err)
	}
}
