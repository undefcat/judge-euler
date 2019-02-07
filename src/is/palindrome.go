package is

func PalindromeInt(n int) bool {
	a := make([]int, 12)

	for n != 0 {
		a = append(a, n%10)
		n /= 10
	}

	return PalindromeIntSlice(a)
}

func PalindromeIntSlice(a []int) bool {
	mid, last := len(a)/2, len(a)-1

	for i := 0; i < mid; i++ {
		if a[i] != a[last-i] {
			return false
		}
	}

	return true
}

func PalindromeString(s string) bool {
	mid, last := len(s)/2, len(s)-1

	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return false
		}
	}

	return true
}