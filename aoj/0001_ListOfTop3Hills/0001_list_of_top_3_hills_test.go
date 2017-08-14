package main

import (
	"reflect"
	"testing"
)

func TestGetTop3(t *testing.T) {
	testCase := []struct {
		desc string
		in   []int
		want []int
	}{
		{
			"Top3を表示できること",
			[]int{1819, 2003, 876, 2840, 1723, 1673, 3776, 2848, 1592, 922},
			[]int{3776, 2848, 2840},
		},
		{
			"同じ高さでも異なる順位になること",
			[]int{100, 200, 300, 400, 500, 600, 700, 800, 900, 900},
			[]int{900, 900, 800},
		},
	}

	for _, test := range testCase {
		t.Run(test.desc, func(t *testing.T) {
			if got := GetTop3(test.in); !reflect.DeepEqual(got, test.want) {
				t.Errorf("GetTop3: got %v want %v", got, test.want)
			}
		})
	}
}
