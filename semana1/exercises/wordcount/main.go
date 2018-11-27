package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	newMap := make(map[string]int)
	strS := strings.Fields(s)
	for i := range strS {
		elem, ok := newMap[strS[i]]
		if !ok {
			newMap[strS[i]] = 1
		} else {
			newMap[strS[i]] = elem + 1
		}
	}
	//return map[string]int{"x": 1}
	return newMap
}

func main() {
	wc := WordCount("I am a new Gopher")
	fmt.Println(wc)
}
