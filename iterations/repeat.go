package iterations

func Repeat(val string, count int) string {
	var character string
	for i := 0; i <= count-1; i++ {
		character += val
	}

	return character
}

var list = []int{1, 3, 3, 4, 5, 6, 7}

func PrintList(val []int) []int {
	var manipulatedList []int
	for _, v := range val {
		manipulatedList = append(manipulatedList, v+1)
	}

	return manipulatedList
}
