package roman

import (
	"testing"
)

func TestToRoman(t *testing.T) {
	testCase := []struct {
		desc string
		in   Arabic
		want Roman
	}{
		{"1=>Iに変換できること", 1, "I"},
	}

	for _, test := range testCase {
		t.Run(test.desc, func(t *testing.T) {
			if got := test.in.ToRoman(); got != test.want {
				t.Errorf("Arabic.ToRoman(): got %v want %v", got, test.want)
			}
		})
	}
}
