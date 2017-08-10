package main

import (
	"testing"
)

func TestMyFunc(t *testing.T) {
	testCase := []struct {
		desc string
		in   int
		want int
	}{
		{
			"与えられた整数nの組み合わせを返すこと1",
			35,
			4,
		},
		{
			"与えられた整数nの組み合わせを返すこと2",
			1,
			4,
		},
		{
			"与えられた整数nの組み合わせを返すこと3",
			2,
			10,
		},
	}

	for _, test := range testCase {
		t.Run(test.desc, func(t *testing.T) {
			if got := GetWays(test.in); got != test.want {
				t.Errorf("GetWays: got %v want %v", got, test.want)
			}
		})
	}
}
