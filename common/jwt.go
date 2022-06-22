package common

import (
	"github.com/golang-jwt/jwt/v4"
	"go-web/model"
	"time"
)

/**

使用 jwt 组件进行token 的配置

*/

var jwtKey = []byte("big-shawn-go-web")

type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

func ReleaseToken(user model.User, ttl time.Duration) (string, error) {

	expireTime := time.Now().Add(ttl * time.Hour)

	claims := &Claims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Big-Shawn",
			Subject:   "User Token",
		},
	}

	// 先是利用clamis 创建一个 token 结构体
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 然后再是创建一个token 字符串
	tokenString, err := token.SignedString(jwtKey)
	/**
	  生成的token 是一个以点号分割的三段式的使用base加密的字符串
	   第一段是 协议头  存储 token 使用的加密方式
			{"alg":"HS256", "type":"JWT"}
	  第二段是  payload base64加密后的值
	  第三段是 前两段加上 jwtkey 进行哈希生成的一个字符串
	*/
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func ParseToken(token string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return jwtToken, claims, err
}
