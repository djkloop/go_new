package main

import (
	"fmt"
	"strconv"
)

func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	fmt.Println(
		converToBin(5),
		converToBin(13),
	)
}
