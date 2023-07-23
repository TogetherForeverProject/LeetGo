package main

func fibonacciMemo(n int, memo map[int]int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
	return memo[n]
}

func fib(n int) int {
	memo := make(map[int]int)
	return fibonacciMemo(n, memo)
}
