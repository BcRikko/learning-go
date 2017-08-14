package main

import (
	"fmt"
	"math"
)

func GetPrimeNumerCount(in int) int {
	primes := make([]bool, in+1)
	primes[0] = false
	primes[1] = false
	for i := 2; i < len(primes); i++ {
		primes[i] = true
	}

	max := math.Floor(math.Sqrt(float64(in)))

	for i := 2; i <= int(max); i++ {
		for j := i * 2; j < len(primes); j += i {
			primes[j] = false
		}
	}

	count := 0
	for i := 0; i < len(primes); i++ {
		if primes[i] == true {
			count++
		}
	}
	return count
}

func main() {
	result := GetPrimeNumerCount(10)
	fmt.Println(result)

}
