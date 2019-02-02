# Double-base palindromes

## Solution1

직관적으로, 배열에 숫자 하나 하나 넣은 뒤 비교한다.(나의 해결법)

## Solution2

수학적인 방법으로 숫자를 조작하여, 원래 숫자를 뒤집은 뒤 비교한다.

```go
func isPalindromic(n, b int) bool {
	reversed := 0
	k := n

	for k > 0 {
		reversed = b*reversed + k%b
		k /= b
	}
	return n == reversed
}
```

## Solution3

10진수 Palindromic 숫자를 1에서부터 차례대로 생성한 다음, 2진수로 Palindromic 한지 검사한다.

이 방법이 가장 빠르다.