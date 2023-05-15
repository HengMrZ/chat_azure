package pkg

import "crypto/rand"

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// RndStr 安全的随机字符串
func RndStr(num ...int) string {
	count := 8
	if len(num) != 0 {
		count = num[0]
		if count <= 0 {
			count = 8
		}
	}
	var bytes = make([]byte, count)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
