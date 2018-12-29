package slices

import (
	"fmt"
	"reflect"
	"testing"
)

var numbers = []int{2, 3, 4, 6, 9, 11}

func TestSum(t *testing.T) {
	t.Run("a collection of 5 numbers", func(t *testing.T) {

		got := Sum(numbers)
		expected := 35

		if got != expected {
			t.Errorf("Expected: '%d', Got: '%d'", expected, got)
		}
	})

	t.Run("a collection of any amount", func(t *testing.T) {
		var numbers = []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("Expected: '%d', Got: '%d'", expected, got)
		}
	})

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want: '%v', Got: '%v'", want, got)
	}
}
