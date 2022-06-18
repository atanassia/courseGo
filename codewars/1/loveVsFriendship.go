package main

import (
	"fmt"
	"strings"
)

func main() {
	count := 0
	for _, i := range strings.Split("attitude", "") {
		fmt.Printf("%s %d\n", i, toCharStr(i))
		count += toCharStr(i)
	}

	fmt.Println(count)
}

func toCharStr(i string) int {
	foo := "abcdefghijklmnopqrstuvwxyz"

	for num, n := range foo {
		if i == string(n) {
			return num + 1
		}
	}
	return 0

}
