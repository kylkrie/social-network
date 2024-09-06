#!/bin/bash

output_file="project_dump.txt"
src_directory="src/lib"

# Clear the output file
>"$output_file"

# Loop through all .ts and .svelte files in src directory, excluding node_modules
find "$src_directory" -type f \( -name "*.ts" -o -name "*.svelte" -o -name "*.css" \) -not -path "*/node_modules/*" | while read file; do
  echo "// $file" >>"$output_file"
  cat "$file" >>"$output_file"
  echo -e "\n" >>"$output_file"
done
