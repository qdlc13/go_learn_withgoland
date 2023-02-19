package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	printFileContents(file)
}
func forever() {
	for {
		fmt.Println("3333")
	}
}
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
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
	currentDir, _ := os.Getwd()
	fmt.Println(currentDir)
	printFile("E:\\git_master\\go_learn_withgoland\\loop\\abc.txt")
	//forever()

	s := `sasd"dsd"
	dddkkks
	ooo o

	p`
	printFileContents(strings.NewReader(s))

}
