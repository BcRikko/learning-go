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
			"n週間後の借金残高を計算できること1",
			5,
			130000,
		},
		{
			"n週間後の借金残高を計算できること2",
			1,
			105000,
		},
		{
			"n週間後の借金残高を計算できること2",
			10,
			168000,
		},
	}

	for _, test := range testCase {
		t.Run(test.desc, func(t *testing.T) {
			if got := GetDept(test.in); got != test.want {
				t.Errorf("MyFunc: got %v want %v", got, test.want)
			}
		})
	}
}
