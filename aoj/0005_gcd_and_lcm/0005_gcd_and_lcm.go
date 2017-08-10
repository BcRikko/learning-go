package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetGCD(x, y int) int {
	if x < y {
		x, y = y, x
	}

	remainder := x % y
	if remainder == 0 {
		return y
	}

	return GetGCD(y, remainder)
}

func GetLCM(x, y, gcd int) int {
	return (x * y) / gcd
}

func GetGCDandLCM(in string) string {
	ins := strings.Split(strings.TrimSpace(in), " ")
	x, _ := strconv.Atoi(ins[0])
	y, _ := strconv.Atoi(ins[1])

	gcm := GetGCD(x, y)
	lcm := GetLCM(x, y, gcm)

	return fmt.Sprintf("%v %v", gcm, lcm)
}

func main() {
	gcm := GetGCD(123, 621)
	lcm := GetLCM(123, 621, gcm)
	fmt.Printf("GCM: %v, LCM: %v", gcm, lcm)
}
