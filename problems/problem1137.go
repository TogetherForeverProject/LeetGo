// N-th Tribonacci Number
//
// The Tribonacci sequence Tn is defined as follows:
// T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
// Given n, return the value of Tn.

package problems

import (
    "fmt"
)

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	t0, t1, t2 := 0, 1, 1

	for i := 3; i <= n; i++ {
		t0, t1, t2 = t1, t2, t0+t1+t2
	}

	return t2
}

// Input: n = 4
func Problem1137() {
	n := 25
	result := tribonacci(n)
	fmt.Printf("Tribonacci number at index %d is: %d\n", n, result)
}

func init() {
    RegisterProblem(1137, Problem1137)
}
