// Leet Code Dynamic Programming
// https://leetcode.com/studyplan/dynamic-programming/

package main

import (
	"flag"
	"github.com/TogetherForeverProject/LeetGo/problems"
)

func main() {
	// Define a command-line flag "problem" to accept the problem number
	problemFlag := flag.Int("problem", 0, "Specify the problem number")
	flag.Parse()

	// Call the RunProblem function from the problems package
	problems.RunProblem(*problemFlag)
}
