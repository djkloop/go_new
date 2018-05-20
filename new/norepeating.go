package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {

	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, char := range []rune(s) {

		if lastI, ok := lastOccurred[char]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[char] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("dsfsfsafdsfssd"))
	fmt.Println(lengthOfNonRepeatingSubStr("你好啊好不好啊啊啊发的看法厚大司考"))
	fmt.Println(lengthOfNonRepeatingSubStr("你好啊你好啊你好不啊"))
}
