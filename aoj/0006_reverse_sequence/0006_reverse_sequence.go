package main

import (
	"fmt"
	"strings"
)

func Reverse(in string) string {
	ret := make([]string, len(in))

	for i, rune := range in {
		ret[len(in)-i-1] = string(rune)
	}

	return strings.Join(ret, "")
}

func main() {
	in := "abc123"
	ans := Reverse(in)
	fmt.Println(ans)
}
