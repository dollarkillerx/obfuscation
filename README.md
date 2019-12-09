# obfuscation
混淆加密 随机匹配

### 安装
`go get github.com/dollarkillerx/obfuscation`
### 简单使用
``` 
func TestDecode(t *testing.T) {
	key := GetRandKey()              // 生成专用key
	log.Println("key: ", key)
	data := "你好亚"
	s, e := Encode(key, []byte(data))  // 编码
	if e != nil {
		panic(e)
	}
	log.Println(s)

 	bytes, e := Decode(key, s)         // 解码
	if e != nil {
		panic(e)
	}
	log.Println(string(bytes)) 
}
```