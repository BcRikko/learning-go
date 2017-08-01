package roman

import (
	"testing"
)

func TestToRoman(t *testing.T) {
	if got, want := ToRoman(1), "I"; got != want {
		t.Errorf("ToRoman(1): got %v want %v", got, want)
	}
}
