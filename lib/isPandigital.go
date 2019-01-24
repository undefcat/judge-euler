package lib

func IsPandigital(n int) bool {
	var digits, count, tmp int

	for n > 0 {
		tmp = digits
		digits |= 1<<uint(n-((n/10)*10)-1)
		if tmp == digits {
			return false
		}
		n /= 10
		count++
	}
	return digits == 1<<uint(count)-1
}