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
			"テストの説明",
			"input",
			"output",
		},
	}

	for _, test := range testCase {
		t.Run(test.desc, func(t *testing.T) {
			if got := MyFunc(); got != test.want {
				t.Errorf("MyFunc: got %v want %v", got, test.want)
			}
		})
	}
}
