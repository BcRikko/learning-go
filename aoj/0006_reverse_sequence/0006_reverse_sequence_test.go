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
			"逆順にした文字列が取得できること",
			"w32nimda",
			"admin23w",
		},
	}

	for _, test := range testCase {
		t.Run(test.desc, func(t *testing.T) {
			if got := Reverse(test.in); got != test.want {
				t.Errorf("MyFunc: got %v want %v", got, test.want)
			}
		})
	}
}
