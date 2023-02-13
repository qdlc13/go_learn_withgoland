package main

import "fmt"

func main() {
	//
	m := map[string]string{
		"name":   "lc",
		"age":    "23",
		"course": "golang",
	}
	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int      //m3 == nil
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v) //不保证顺序 是无序的 每次打印顺序不一定相同
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	cou, ok := m["cou"]  //cou没定义过所以可用:=
	fmt.Println(cou, ok) //空串
	if NValue, ok := m["xxx"]; ok {
		fmt.Println(m[NValue])
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"] //两者都定义过只能 =
	fmt.Println(name, ok)

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

func lengthOfNonRepeatingSubStr(s string) int {
	maxLength := 0
	start := 0
	lastOccured := make(map[rune]int)
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastOccured[ch] + 1
		}
		if maxLength < i-start+1 {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}
