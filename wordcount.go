package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	split := strings.Fields(s)

	m := make(map[string]int)

	for i := range split {
		elem, ok := m[split[i]] // split[i]에 해당하는 string 값이 있으면 elem을 가져옴

		if ok == true { //있는게 또 나온 경우 +1을 해줌
			m[split[i]] = elem + 1
		} else if ok == false {
			m[split[i]] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
