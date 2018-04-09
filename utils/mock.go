package utils

import (
	"crypto/rand"
	mr "math/rand"
	"fmt"
)

// genToken 生成随机字符串
func GenToken(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// 获取随机数字
func GenNub(n int) string {
	const alphanum = "0123456789"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// 获取随机字符串
func GenStr(n int,alphanum string) string {
	if alphanum == ""{
		alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// 获取布尔值 2为随机生成
func GenBool(n int) bool {
	if n == 2{
		const alphanum = "01"
		var bytes = make([]byte, n)
		rand.Read(bytes)
		for i, b := range bytes {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		}
		if string(bytes) == "1"{
			return true
		} else {
			return false
		}
	}
	if n == 0{
		return false
	}else {
		return true
	}
}

// 获取某一段数字
func GetMNub(start,end int) int {
	for{
		ret := mr.Intn(end)
		if ret >= start{
			return ret
		}
	}
}

// 获取随机小数
func GenFloat(n,f int) string {
	const alphanum = "0123456789"
	var bytesn = make([]byte, n)
	var bytesf = make([]byte, f)
	rand.Read(bytesn)
	rand.Read(bytesf)
	for i, b := range bytesn {
		bytesn[i] = alphanum[b%byte(len(alphanum))]
	}
	for i, b := range bytesf {
		bytesf[i] = alphanum[b%byte(len(alphanum))]
	}
	return fmt.Sprintf("%s.%s",string(bytesn),string(bytesf))
}