#!/bin/bash


# dont do this

# Ensure the script is executable
# chmod +x run_loop.sh

# Path to your main.go file
MAIN_GO_PATH="main.go"

# Loop from 0 to 10000
for X in {0..10000}
do
    # Set the environment variable
    export X=$X
    
    # Run the Go program
    go run "$MAIN_GO_PATH"
    
    # Optionally, you can log the current value of X
    echo "Ran main.go with X=$X"
done
