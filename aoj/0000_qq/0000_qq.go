package main

import (
	"fmt"
)

func QQ() string {
	var ret string
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			ret += fmt.Sprintf("%vx%v=%v\n", i, j, i*j)
		}
	}

	return ret
}

func main() {
	ret := QQ()
	fmt.Print(ret)
}
