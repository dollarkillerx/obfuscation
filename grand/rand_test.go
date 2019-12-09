/**
 * @Author: DollarKillerX
 * @Description: rand_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午6:03 2019/12/9
 */
package grand

import (
	"fmt"
	"testing"
)

func TestRandAbcAc(t *testing.T) {
	s := "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm"
	for _, i := range s {
		fmt.Printf("\"%s\",", string(i))
	}
}

//func TestRandLenNum(t *testing.T) {
//	num := RandLenNum(32)
//	log.Println(num)
//}
