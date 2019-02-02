package is

func Prime(n int) bool {
	switch {
	case n == 1:
		return false

	case n == 2, n == 3:
		return true

	case n%2 == 0, n%3 == 0:
		return false
	}

	i, w := 5, 2

	for i*i <= n {
		if n%i == 0 {
			return false
		}

		i, w = i+w, 6-w
	}
	return true
}