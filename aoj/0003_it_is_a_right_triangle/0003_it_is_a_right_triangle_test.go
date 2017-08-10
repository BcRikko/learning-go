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
			"直角三角形の判定ができること",
			`3
4 3 5
4 3 6
8 8 8`,
			`YES
NO
NO`,
		},
	}

	for _, test := range testCase {
		t.Run(test.desc, func(t *testing.T) {
			if got := IsRightTriangle(test.in); got != test.want {
				t.Errorf("IsRightTriangle: got %v want %v", got, test.want)
			}
		})
	}
}
