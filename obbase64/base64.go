/**
 * @Author: DollarKillerX
 * @Description: 实现基础base64
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:47 2019/12/9
 */
package obbase64

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// 获取md5
func Md5Encode(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 获取sha1
func Sha1Encode(str string) string {
	data := []byte(str)
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

// 获取sha256
func Sha256Encode(str string) string {
	sum256 := sha256.Sum256([]byte(str))
	s := hex.EncodeToString(sum256[:])
	return s
}

// Base64编码
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64解码
func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// Base64URL编码
func Base64URLEncode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// Base64URL解码
func Base64URLDecode(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}
