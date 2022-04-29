package ms

import (
	"testing"
)

func TestParseBool(t *testing.T) {
	tf := ParseBool("Yes")
	if !tf {
		t.Errorf("Error ParseBool failed\n")
	}
	tf = ParseBool("")
	if tf {
		t.Errorf("Error ParseBool failed\n")
	}
}
