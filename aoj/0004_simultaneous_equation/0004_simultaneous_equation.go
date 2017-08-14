package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseFloat(in string) []float64 {
	var ret []float64
	for _, v := range strings.Split(strings.TrimSpace(in), " ") {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			ret = append(ret, f)
		} else {
			ret = append(ret, 0)
		}
	}
	return ret
}

func CalcEquation(in string) string {
	// 1. ax + by = c
	// 2. dx + ey = f
	// 1. (a * d)x + (b * d)y = c * d
	// 2. (d * a)x + (e * a)y = f * a
	// (b * d) - (e * a) y = (c * d) - (f * a)
	// y = ((c * d) - (f * a)) / ((b * d) - (e * a))
	// ax + by = c
	// x = (c - by) / a
	ins := ParseFloat(in)
	a, b, c, d, e, f := ins[0], ins[1], ins[2], ins[3], ins[4], ins[5]

	y := ((c * d) - (f * a)) / ((b * d) - (e * a))
	x := (c - (b * y)) / a

	return fmt.Sprintf("%.3f %.3f", x, y)
}

func main() {
	in := "1 2 3 4 5 6"
	ans := CalcEquation(in)
	fmt.Println(ans)
}
