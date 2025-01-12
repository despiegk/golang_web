#!/bin/bash

# Make script executable from anywhere
cd "$(dirname "$0")/heroweb"

# Check if we want to run in development mode with live reload
if [ "$1" = "--dev" ]; then
    echo "Starting in development mode with live reload..."
    make watch
else
    echo "Starting application..."
    make run
fi
