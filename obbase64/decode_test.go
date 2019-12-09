/**
 * @Author: DollarKillerX
 * @Description: decode_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午6:26 2019/12/9
 */
package obbase64

import (
	"log"
	"strings"
	"testing"
)

func TestGetRandKey(t *testing.T) {
	key := GetRandKey()
	log.Println(key)

	split := strings.Split(key, ",")
	log.Println(split)
}

func TestBase64Decode(t *testing.T) {
	data := "我是你的爸爸"
	encode := Base64Encode([]byte(data))
	log.Println(encode)
	bytes, e := Base64Decode(encode)
	if e != nil {
		panic(e)
	}
	log.Println(string(bytes))
}

func TestDecode(t *testing.T) {
	key := GetRandKey()
	log.Println("key: ", key)
	data := "你好亚"
	s, e := Encode(key, []byte(data))
	if e != nil {
		panic(e)
	}
	log.Println(s)

	bytes, e := Decode(key, s)
	if e != nil {
		panic(e)
	}
	log.Println(string(bytes))
}
