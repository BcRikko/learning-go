package main

import (
	"testing"
)

func TestMyFunc(t *testing.T) {
	testCase := []struct {
		desc string
		in   string
		want string
	}{
		{
			"最大公約数と最小公倍数が取得できること1",
			"8 6",
			"2 24",
		},
		{
			"最大公約数と最小公倍数が取得できること2",
			"50000000 30000000",
			"10000000 150000000",
		},
		{
			"最大公約数と最小公倍数が取得できること3",
			"123 621",
			"3 25461",
		},
	}

	for _, test := range testCase {
		t.Run(test.desc, func(t *testing.T) {
			if got := GetGCDandLCM(test.in); got != test.want {
				t.Errorf("GetGCDandLCM: got %v want %v", got, test.want)
			}
		})
	}
}
