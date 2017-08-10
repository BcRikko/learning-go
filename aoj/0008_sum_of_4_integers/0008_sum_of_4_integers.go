package main

import (
	"fmt"
)

func GetWays(n int) int {
	count := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				for l := 0; l < 10; l++ {
					if i+j+k+l == n {
						count++
					}
				}
			}
		}
	}

	return count
}

func main() {
	result := GetWays(2)
	fmt.Println(result)
}
