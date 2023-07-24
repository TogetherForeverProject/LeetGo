#!/bin/bash

print_error() {
    local text="$1"
    local red
    red=$(tput setaf 1)
    local reset
    reset=$(tput sgr0)

    echo "[${red}Error${reset}] $text"
}

get_first_comment() {
    local filename="$1"
    grep -m 1 -E '^\s*//' "$filename" | sed 's/^\s*\/\/\s*//'
}

# Check if the number argument is provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <problem_number>"
    exst 1
fi

# Get the problem number from the first argument
problem_number=$1
bold=$(tput bold)

cyan=$(tput setaf 6)
reset=$(tput sgr0)

file="problems/problem$problem_number.go"
# Check if the problem file exists in the problems folder
if [ -f "$file" ]; then
    # If the file exists, run the main.go with the specified problem number
    problem_title=$(get_first_comment "$file")
    echo "${bold}Problem $problem_number ($problem_title)"
    echo "──────────────────────────────"
    output_go=$(go run main.go -problem="$problem_number")

    # Print both outputs on the same line using printf
    printf "[${cyan}Output${reset}] %s\n" "$output_go"
else
    # If the file does not exist, display an error message
    print_error "Problem $problem_number does not exist."
fi
