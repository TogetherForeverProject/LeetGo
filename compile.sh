#!/bin/bash

print_error() {
    local text="$1"
    local red
    red=$(tput setaf 1)
    local reset
    reset=$(tput sgr0)

    echo "[${red}Error${reset}] $text"
}

# Check if the number argument is provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <problem_number>"
    exit 1
fi

# Get the problem number from the first argument
problem_number=$1
bold=$(tput bold)

cyan=$(tput setaf 6)
reset=$(tput sgr0)
# Check if the problem file exists in the problems folder
if [ -f "problems/problem$problem_number.go" ]; then
    # If the file exists, run the main.go with the specified problem number
    echo "${bold}Problem $problem_number"
    echo "──────────────────────────────"
    output_go=$(go run main.go -problem="$problem_number")

    # Print both outputs on the same line using printf
    printf "[${cyan}Output${reset}] %s\n" "$output_go"
else
    # If the file does not exist, display an error message
    print_error "Problem $problem_number does not exist."
fi
