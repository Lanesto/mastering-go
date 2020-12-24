// Package documentme is for testing Go documentation check
// no special functionalities
package documentme

// Pie is global variable
// -
const Pie = 3.1415912

// S1 get length of given string
// use range to loop over string
func S1(s string) int {
	if s == "" {
		return 0
	}
	n := 0
	for range s {
		n++
	}
	return n
}

// F1 returns twice of given integer
// Double() would be better naming
func F1(n int) int {
	return 2 * n
}
