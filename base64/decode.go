/**
 * @Author: DollarKillerX
 * @Description: 混淆base64加密
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:52 2019/12/9
 */
package base64

import (
	"log"
	"strconv"
	"strings"

	"github.com/dollarkillerx/obfuscation/grand"
)

// 混淆原理
// 通过key 对base64的特定位置 增加原子 达到混淆的效果  32位密钥 每2位是一组

// 获得一个随机的key
func GetRandKey() string {
	result := ""
	for i := 1; i <= 32; i++ {
		if i%2 == 0 && i != 32 {
			result += strconv.Itoa(grand.RandNum()) + ","
		} else {
			result += strconv.Itoa(grand.RandNum())
		}
	}
	return result
}

// 解码
func Decode(key string, data string) ([]byte, error) {
	keys := strings.Split(key, ",")
	for _, k := range keys {
		i, e := strconv.Atoi(k)
		if e != nil {
			return nil, e
		}
		if i%2 == 0 {
			data = data[:len(data)-1]
		} else {
			data = data[1:]
		}
	}
	return Base64Decode(data)
}

// 编码
func Encode(key string, data []byte) (string, error) {
	keys := strings.Split(key, ",")
	encode := Base64Encode(data)
	log.Println(encode)

	for _, k := range keys {
		i, e := strconv.Atoi(k)
		if e != nil {
			return "", e
		}
		if i%2 == 0 {
			encode += grand.RandAbc()
		} else {
			encode = grand.RandAbc() + encode
		}
	}
	return encode, nil
}
