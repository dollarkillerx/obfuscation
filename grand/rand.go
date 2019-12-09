/**
 * @Author: DollarKillerX
 * @Description: grand.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午6:01 2019/12/9
 */
package grand

import (
	"math/rand"
	"time"
)

var abc = []string{
	"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P", "A", "S", "D", "F", "G", "H", "J", "K", "L", "Z", "X", "C", "V", "B", "N", "M", "q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
}

// 获取随机字母 和数字
func RandAbc() string {
	rand.Seed(time.Now().UnixNano())
	return abc[rand.Intn(len(abc))]
}

// 获取随机数
func RandNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}
