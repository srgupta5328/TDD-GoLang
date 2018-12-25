package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected: '%s', got: '%s'", expected, repeated)
	}
}

//When comparing lists make sure that it equals nil for test
func TestPrintLnist(t *testing.T) {
	got := PrintList(list)
	expected := []int{2, 4, 4, 5, 6, 7, 8}

	if got == nil {
		t.Errorf("expected: '%d', got: '%d'", expected, got)
	}

}
