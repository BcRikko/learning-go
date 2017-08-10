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
			"連立方程式の解が求められること1",
			"1 2 3 4 5 6",
			"-1.000 2.000",
		},
		{
			"連立方程式の解が求められること2",
			"2 -1 -2 -1 -1 -5",
			"1.000 4.000",
		},
		{
			"連立方程式の解が求められること3",
			"2 -1 -3 1 -1 -3",
			"0.000 3.000",
		},
		{
			"連立方程式の解が求められること4",
			"2 -1 -3 -9 9 27",
			"0.000 3.000",
		},
	}

	for _, test := range testCase {
		t.Run(test.desc, func(t *testing.T) {
			if got := CalcEquation(test.in); got != test.want {
				t.Errorf("MyFunc: got %v want %v", got, test.want)
			}
		})
	}
}
