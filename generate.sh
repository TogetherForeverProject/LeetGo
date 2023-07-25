#!/bin/bash

# Define variables in a table-like format
variables=$(cat <<EOF
title=LeetGo
easy=<span style="color:rgb(0 184 163);">Easy</span>
medium=<span style="color:rgb(255 192 30);">Medium</span>
hard=<span style="color:rgb(255 55 95);">Hard</span>
EOF
)

# Read the template.md content
content=$(<template.md)

# Function to replace placeholders with their corresponding variable values
replace_placeholders() {
    while read -r line; do
        local var_name=${line%%=*}
        local var_value=${line#*=}
        content=${content//"{{${var_name}}}"/"${var_value}"}
    done <<< "$variables"

    # Replace [[number]] with [problem$number.go](./problems/problem$number.go)
    content=$(echo "$content" | sed -E 's/\[\[([0-9]+)\]\]/\[problem\1.go\](\.\/problems\/problem\1.go\)/g')
}

# Call the function to replace placeholders
replace_placeholders

# Save the modified content to README.md
echo "$content" > README.md
