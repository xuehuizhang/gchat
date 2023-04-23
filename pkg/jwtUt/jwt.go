package jwtUt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"time"
)

var JwtSecret = []byte("1234567890")

type Claims struct {
	Username string `json:"username"`
	Id       int64  `json:"id"`
	Password string `json:"password"`
	jwt.StandardClaims
	Type int `json:"type"` //0:用户名密码登陆， 1：外部手机号登陆
}

func GenerateToken(id int64, username, password string, loginType int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		username,
		id,
		EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "https://gchat.com",
		},
		loginType,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func GetIdFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)

			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}

func GetUidFromToken(token string) int64 {
	c, err := ParseToken(token)

	if err != nil {
		return -1
	}

	return c.Id
}

func GetClaimsFromToken(token string) (*Claims, error) {
	c, err := ParseToken(token)

	if err != nil {
		return nil, err
	}
	return c, nil
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
