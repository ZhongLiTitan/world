package world

import "testing"

func ATest(t *testing.T) {
	want := "a"
	if got := A(); got != want {
		t.Errorf("a = %q,want %q", got, want)
	}
}
