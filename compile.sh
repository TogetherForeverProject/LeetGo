#!/bin/bash

# Check if the number argument is provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <problem_number>"
    exit 1
fi

# Get the problem number from the first argument
problem_number=$1

# Check if the problem file exists in the problems folder
if [ -f "problems/problem$problem_number.go" ]; then
    # If the file exists, run the main.go with the specified problem number
    echo "Problem $problem_number"
    echo "===================="
    go run main.go -problem="$problem_number"
else
    # If the file does not exist, display an error message
    echo "Problem file for problem $problem_number does not exist."
fi
