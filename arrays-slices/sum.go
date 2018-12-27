package slices

var numbers = []int{2, 3, 4, 6, 9, 11}

func Sum(nums []int) int {
	var total int
	for _, v := range nums {
		total += v
	}
	return total
}
