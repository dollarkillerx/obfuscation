/**
 * @Author: DollarKillerX
 * @Description: 压缩字符串  高效网络传输(注意:如果数据比较小 压缩反之会便大)
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:04 2019/12/9
 */
package zipstr

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

func Zip(str string) string {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(str)); err != nil {
		return ""
	}
	if err := gz.Flush(); err != nil {
		return ""
	}
	if err := gz.Close(); err != nil {
		return ""
	}
	strc := base64.StdEncoding.EncodeToString(b.Bytes())
	return strc
}

func Unzip(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	rdata := bytes.NewReader(data)
	rc, err := gzip.NewReader(rdata)
	if err != nil {
		return ""
	}
	all, err := ioutil.ReadAll(rc)
	if err != nil {
		return ""
	}
	return string(all)
}
