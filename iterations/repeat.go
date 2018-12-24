package iterations

func Repeat(val string) string {
	var character string
	for i := 0; i <= 5; i++ {
		character = val + character
	}
	return character
}
