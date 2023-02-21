package main

import "fmt"

// 优化map提前开好空间 用空间换时间
var lastOccured = make([]int, 0xffff) //65535
func lengthOfNonRepeatingSubStr(s string) int {
	//发现map和rune是耗时间的主要部分
	//优化map提前开好空间 用空间换时间
	for i := range lastOccured {
		lastOccured[i] = -1
	}
	maxLength := 0
	start := 0
	//lastOccured := make(map[rune]int)
	for i, ch := range []rune(s) {
		if lastI := lastOccured[ch]; lastI != -1 && lastI >= start {
			start = lastOccured[ch] + 1
		}
		if maxLength < i-start+1 {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}
func main() {
	//寻找最长不含有重复字符的子串 nonrepeating
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("v"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("我是初学者"))
	fmt.Println(lengthOfNonRepeatingSubStr("我eeee是初学者"))

}
