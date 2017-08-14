package main

import (
	"fmt"
	"sort"
)

func GetTop3(inputs []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(inputs)))
	return inputs[:3]
}

func main() {
	ints := []int{5, 2, 4, 3, 1}
	fmt.Println(ints)

	sorted := GetTop3(ints)
	fmt.Println(sorted)
}
