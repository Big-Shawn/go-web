package util

import (
	"math/rand"
	"time"
)

// RandomString  生成一个随机的字符串
func RandomString(n int) string {

	var strings = []byte("asdfaqetryuioplamznbvASDFGHQWERTYUIOPLZMNCBV")

	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = strings[rand.Intn(len(strings))]
	}

	return string(result)
}
