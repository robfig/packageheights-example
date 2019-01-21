package b

import "testing"

func TestB(t *testing.T) {
	if Name() != "b" {
		t.Error("unexpected name")
	}
}
