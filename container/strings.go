package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱大海啊"
	fmt.Println(len(s))           //18字节数
	fmt.Printf("%s\n", s)         //Yes我爱大海啊
	fmt.Printf("%X\n", s)         //596573E68891E788B1E5A4A7E6B5B7E5958A
	fmt.Printf("%s\n", []byte(s)) //Yes我爱大海啊
	fmt.Printf("%X\n", []byte(s)) //596573E68891E788B1E5A4A7E6B5B7E5958A

	//[]byte(s) 可以获得字节
	//这是UTF-8编码 可变长 英文一字节中文三字节
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) //59 65 73 E6 88 91 E7 88 B1 E5 A4 A7 E6 B5 B7 E5 95 8A
	}
	fmt.Println()
	for i, b := range s { //b is a rune   rune等价于int32  这是 unicode编码
		fmt.Printf("(%d %X) ", i, b) //(0 59) (1 65) (2 73) (3 6211) (6 7231) (9 5927) (12 6D77) (15 554A)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) //8字符数 直接len(s)是字节数

	//DecodeRune 解压缩 p 中的第一个 UTF-8 编码并返回符文及其宽度(以字节为单位)。如果 p 为空，则返回 (RuneError, 0)。
	//否则，如果编码无效，则返回 (RuneError, 1)。对于正确的、非空的 UTF-8，两者都是不可能的结果。
	bytes := []byte(s)
	fmt.Println(bytes)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	} //Yes我爱大海啊
	fmt.Println()

	for i, ch := range []rune(s) { //decode 一个字符占四个字节重新开辟空间 不是用别的方法解释内存
		fmt.Printf("(%d %c)", i, ch)
	} //(0 Y)(1 e)(2 s)(3 我)(4 爱)(5 大)(6 海)(7 啊)
	fmt.Println()

}
