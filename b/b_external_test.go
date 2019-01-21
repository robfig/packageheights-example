package b_test

import "testing"
import "packageheights/a"
import "packageheights/b"

func TestAB(t *testing.T) {
	if a.Name() != "ab" {
		t.Error("unexpected name")
	}
}

func TestB(t *testing.T) {
	if b.Name() != "b" {
		t.Error("unexpected name")
	}
}
