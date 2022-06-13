package util

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
)

var (
	signKey *rsa.PrivateKey
)

type JWTClaims struct {
	*jwt.RegisteredClaims
	TokenType string
	payload   interface{}
}

// 这里怎么找不到一个合适的数据类型来存放有效的时间
// 要使用Jwt自带的ttl生成函数来设置过期时间
// 另外 这个token 需要写成中间件的形式来进行处理

func SetToken(ttl *jwt.NumericDate, method string, payload ...interface{}) (string, error) {
	t := jwt.New(jwt.GetSigningMethod(method))
	t.Claims = &JWTClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: ttl,
		},
		"Login Token",
		payload,
	}

	return t.SignedString(signKey)
}
