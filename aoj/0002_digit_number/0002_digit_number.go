package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Add(in string) string {
	var ret []string

	for _, ln := range strings.Split(in, "\n") {
		ins := strings.Split(strings.TrimSpace(ln), " ")

		x, _ := strconv.Atoi(ins[0])
		y, _ := strconv.Atoi(ins[1])

		ret = append(ret, fmt.Sprintf("%v", x+y))
	}

	return strings.Join(ret, "\n")
}

func main() {
	in := `1 2
3 4 
5 6`
	fmt.Println(in)

	add := Add(in)
	fmt.Println(add)
}
