package main

import (
	"testing"
)

func TestGetPrimeNumerCount(t *testing.T) {
	testCase := []struct {
		desc string
		in   int
		want int
	}{
		{
			"整数n以下の素数の個数が取得できること1",
			10,
			4,
		},
		{
			"整数n以下の素数の個数が取得できること2",
			3,
			2,
		},
		{
			"整数n以下の素数の個数が取得できること3",
			11,
			5,
		},
	}

	for _, test := range testCase {
		t.Run(test.desc, func(t *testing.T) {
			if got := GetPrimeNumerCount(test.in); got != test.want {
				t.Errorf("GetPrimeNumerCount: got %v want %v", got, test.want)
			}
		})
	}
}
