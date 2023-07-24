package problems

import (
    "flag"
    "fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// ProblemFunction represents the signature of the problem functions
type ProblemFunction func()

// Define a map that maps problem numbers to their corresponding functions
var problemMap = make(map[int]ProblemFunction)

// RegisterProblem registers a problem function with the problemMap
func RegisterProblem(problemNumber int, problemFunc ProblemFunction) {
	problemMap[problemNumber] = problemFunc
}

// RunProblem runs the problem function for the given problem number
func RunProblem(problemNumber int) {
	// Check if the function exists in the map
	if problemFunc, ok := problemMap[problemNumber]; ok {
		// Call the function
		problemFunc()
	} else {
		fmt.Println("Invalid problem number.")
		flag.PrintDefaults()
	}
}

