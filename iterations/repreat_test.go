package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected: '%s', got: '%s'", expected, repeated)
	}
}
