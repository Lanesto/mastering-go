package ex

func f1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return f1(n-1) + f1(n-2)
}

func s1(s string) int {
	return len(s)
}
