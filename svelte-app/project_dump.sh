#!/bin/bash

output_file="project_dump.txt"
src_directory="."

# Define arrays for file types to include and directories to ignore
include_types=("*.ts" "*.js" "*.svelte" "*.css")
ignore_dirs=("node_modules" "build" ".svelte-kit")

# Clear the output file
>"$output_file"

# Construct the find command
find_cmd="find \"$src_directory\""

# Add ignore patterns
for dir in "${ignore_dirs[@]}"; do
  find_cmd+=" -name $dir -prune -o"
done

# Add include patterns
find_cmd+=" -type f \("
for type in "${include_types[@]}"; do
  find_cmd+=" -name \"$type\" -o"
done

# Remove the last " -o" and close the parenthesis
find_cmd="${find_cmd% -o} \)"

# Add print action
find_cmd+=" -print"

# Execute the find command and process files
eval "$find_cmd" | while read -r file; do
  echo "// $file" >>"$output_file"
  cat "$file" | sed '/^\s*$/d' >>"$output_file"
  echo >>"$output_file"
done

# Remove empty lines from the entire output file
sed -i '/^\s*$/d' "$output_file"
