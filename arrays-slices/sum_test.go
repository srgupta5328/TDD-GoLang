package slices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum(numbers)
	expected := 35

	if got != expected {
		t.Errorf("Expected: '%d', Got: '%d'", expected, got)
	}
}

func ExampleSum() {
	result := Sum(numbers)
	fmt.Println(result)

	//Output: 35
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
