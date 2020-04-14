package helper

// Inverse takes a stack string and creates the inverse of it
// Given "--++--" Inverse returns "++--++"
func Inverse(stack string) string {
	var result string
	for _, s := range stack {
		if s == '-' {
			result += "+"
		}
		if s == '+' {
			result += "-"
		}
	}
	return result
}

// Reverse returns the string parameter in reverse order
func Reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
