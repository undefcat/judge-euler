# Pandigital multiples

## Solution

`isPandigital`은 `problem32`에서 이미 사용한 바 있다. 일단 이 함수를 사용하기로 한다.

문제에서 요구하는 답의 `upper-bound`는 4자리수의 곱이다. `n > 1`인 경우, 즉 `n * 1` `n * 2`까지는 해야 하므로 5자리를 하게되면 2를 곱한 상황에서 10자리를 넘어가게된다. 따라서 4자리까지만 하면 된다.

그 뒤에는 그냥 단순 구현이다.