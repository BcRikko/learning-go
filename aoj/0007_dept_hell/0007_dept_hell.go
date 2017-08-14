package main

import (
	"fmt"
	"math"
)

func GetDept(week int) int {
	const dept float64 = 100000
	const rate float64 = 1.05 // 5%
	ret := dept

	for i := 0; i < week; i++ {
		// 1000円未満切り上げ
		ret = math.Ceil((ret*rate)/1000) * 1000
	}

	return int(ret)
}

func main() {
	dept := GetDept(5)
	fmt.Println(dept)
}
