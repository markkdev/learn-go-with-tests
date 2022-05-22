package iteration

// Repeat takes a string character and int num to repeat a character num times
func Repeat(character string, num int) (repeated string) {
	for i := 0; i < num; i++ {
		repeated += character
	}
	return repeated
}