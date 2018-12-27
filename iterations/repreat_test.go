package iterations

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	result := Repeat("a", 6)
	fmt.Println(result)

	//Output: aaaaaa
}
func TestRepeat(t *testing.T) {
	t.Run("Testing the repeat function with 'a'", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"

		if repeated != expected {
			t.Errorf("expected: '%s', got: '%s'", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

//When comparing lists make sure that it equals nil for test
func TestPrintList(t *testing.T) {
	t.Run("Testing with a list of integers", func(t *testing.T) {
		got := PrintList(list)
		expected := []int{2, 4, 4, 5, 6, 7, 8}

		if got == nil {
			t.Errorf("expected: '%d', got: '%d'", expected, got)
		}
	})

}
