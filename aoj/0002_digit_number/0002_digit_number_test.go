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
			"2つの整数の和が求められること",
			`5 7
1 99
1000 999`,
			`12
100
1999`,
		},
		{
			"前後の余分なスペースは無視されること",
			` 1 2   
3 4
   5 6`,
			`3
7
11`,
		},
	}

	for _, test := range testCase {
		t.Run(test.desc, func(t *testing.T) {
			if got := Add(test.in); got != test.want {
				t.Errorf("Add: got %v want %v", got, test.want)
			}
		})
	}
}
