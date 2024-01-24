package iteration

// Repeat a given string a given number of times
func Repeat(str string, rep int) string {
	repeated := ""
	for i := 0; i < rep; i++ {
		repeated += str
	}
	return repeated
}
