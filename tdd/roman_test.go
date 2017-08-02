package roman

import (
	"testing"
)

func TestToRoman(t *testing.T) {
	if got, want := Arabic(1).ToRoman(), Roman("I"); got != want {
		t.Errorf("Arabic(1).ToRoman(): got %v want %v", got, want)
	}
}
