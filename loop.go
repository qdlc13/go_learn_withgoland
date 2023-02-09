package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		// strconv.Itoa() 把参数转成字符型
		result = strconv.Itoa(lsb) + result
	}
	if result == "" {
		result = "0"
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	//省略其他条件只剩结束条件类似while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		//装二进制
		convertToBin(5), //101
		convertToBin(13),
		convertToBin(0)) //1101

	printFile("abc.txt")

}
