#!/bin/bash

output_file="project_dump.txt"
src_directory="."

# Clear the output file
>"$output_file"

# Loop through all files in src directory excluding index.ts
#find "$src_directory" \( -name "*.go" -o -name "*.sql" \) -type f | while read file; do
find "$src_directory" \( -name "*.go" \) -type f | while read file; do
  echo "// $file" >>"$output_file"
  cat "$file" >>"$output_file"
  echo -e "\n" >>"$output_file"
done
