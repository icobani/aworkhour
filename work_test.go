package work

import "testing"

func TestSimple(t *testing.T) {
	b := 2
	Add(3, &b)
	if b != 5 {
		t.Error("5 değil")
	}
}
