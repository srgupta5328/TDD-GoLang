package iterations

func Repeat(val string) string {
	var character string
	for i := 0; i <= 5; i++ {
		character = val + character
	}
	return character
}

var list = []int{1, 3, 3, 4, 5, 6, 7}

func PrintList(val []int) []int {
	var manipulatedList []int
	for _, v := range list {
		manipulatedList = append(manipulatedList, v+1)
	}

	return manipulatedList
}
