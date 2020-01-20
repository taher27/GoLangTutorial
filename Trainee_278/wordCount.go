package main

import (
	"fmt"
	// "code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)

	for _, v := range strings.Fields(s) {
		if _, ok := m[v]; ok == true {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	return m
}

func main() {
	n := make(map[string]int)
	n = WordCount("hey I am Taher and I am a boy")
	fmt.Println(n)
	// wc.Test(WordCount)
}
