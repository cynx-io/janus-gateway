#!/bin/bash

# Get the directory relative to the script location
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
DIR="$SCRIPT_DIR/../web"

# Function to clean a single file
clean_file() {
    local file_path="$1"

    # Create a temporary file
    local temp_file=$(mktemp)

    # Process the file with sed
    sed -E '
        # Remove the import line
        /^import\s*\{\s*file_buf_validate_validate\s*\}\s*from\s*["'"'"'][^"'"'"']*["'"'"'];\s*$/d
        # Remove from array references
        s/,\s*file_buf_validate_validate//g
        s/file_buf_validate_validate,\s*//g
        s/\[\s*file_buf_validate_validate\s*\]/[]/g
    ' "$file_path" > "$temp_file"

    # Replace original file with cleaned content
    mv "$temp_file" "$file_path"

    echo "✔ Cleaned: $file_path"
}

# Function to walk directory recursively
walk_directory() {
    local dir="$1"

    # Find all .ts files recursively and process them
    find "$dir" -type f -name "*.ts" | while read -r file; do
        clean_file "$file"
    done
}

# Check if proto directory exists
if [ ! -d "$DIR" ]; then
    echo "Error: Directory $DIR does not exist"
    exit 1
fi

# Start the cleaning process
echo "Cleaning TypeScript files in $DIR..."
walk_directory "$DIR"
echo "Done!"