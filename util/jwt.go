package util

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var mySecret = []byte("this is the jwt secret from go-web")

// 这里的引用和结构体有啥区别 传的不都是内存地址吗？

type JWTClaims struct {
	*jwt.RegisteredClaims
	payload interface{}
}

// 这里怎么找不到一个合适的数据类型来存放有效的时间
// 要使用Jwt自带的ttl生成函数来设置过期时间
// 另外 这个token 需要写成中间件的形式来进行处理

/**


使用jwt 生成token

1. 封装 token 生成函数，包含以下参数
	有效期
    加密内容
2. 生成claim 实例
	使用 jwt 内置的 RegisteredClaims 结构，并填充对应的信息
3. 生成 token 对象

4. 返回token 字符串

*/

func SetToken(ttl time.Duration, payload ...interface{}) (string, error) {
	claim := JWTClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		payload: payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(mySecret)
}
