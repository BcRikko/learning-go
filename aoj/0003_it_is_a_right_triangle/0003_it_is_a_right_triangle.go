package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func mapping(items []string, f func(string) float64) []float64 {
	ret := make([]float64, len(items))
	for i, v := range items {
		ret[i] = f(v)
	}
	return ret
}

func IsRightTriangle(in string) string {
	ins := strings.Split(in, "\n")
	n, _ := strconv.Atoi(ins[0])

	var ret []string
	for i := 1; i <= n; i++ {
		ln := strings.Split(strings.TrimSpace(ins[i]), " ")

		sides := mapping(ln, func(v string) float64 {
			n, _ := strconv.ParseFloat(v, 64)
			return math.Pow(n, 2)
		})

		if sides[0]+sides[1] == sides[2] ||
			sides[1]+sides[2] == sides[0] ||
			sides[2]+sides[0] == sides[1] {

			ret = append(ret, "YES")
		} else {
			ret = append(ret, "NO")
		}
	}

	return strings.Join(ret, "\n")
}

func main() {
	in := `3
4 3 5
4 3 6
8 8 8`
	result := IsRightTriangle(in)
	fmt.Println(result)
}
