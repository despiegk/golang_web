#!/bin/bash

# Make script executable from anywhere
cd "$(dirname "$0")"

echo "Installing dependencies..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go first."
    exit 1
fi

# Install go-blueprint if not already installed
if ! command -v go-blueprint &> /dev/null; then
    echo "Installing go-blueprint..."
    go install github.com/melkeydev/go-blueprint@latest
fi

# Install air for live reload if not already installed
if ! command -v air &> /dev/null; then
    echo "Installing air for live reload..."
    go install github.com/air-verse/air@latest
fi

# Install templ if not already installed
if ! command -v templ &> /dev/null; then
    echo "Installing templ..."
    go install github.com/a-h/templ/cmd/templ@latest
    GOPATH=$(go env GOPATH)
    echo "export PATH=\$PATH:$GOPATH/bin" >> ~/.zshrc
    export PATH=$PATH:$GOPATH/bin
fi

# Check if Homebrew is installed (for macOS)
if command -v brew &> /dev/null; then
    # Install tailwindcss using Homebrew if not already installed
    if ! command -v tailwindcss &> /dev/null; then
        echo "Installing tailwindcss..."
        brew install tailwindcss
    fi
else
    echo "Warning: Homebrew not found. Please install tailwindcss manually."
fi

# Install Go dependencies
echo "Installing Go dependencies..."
cd heroweb
go mod download
go mod tidy

# Generate templ files and build CSS
echo "Generating template files and building CSS..."
templ generate
tailwindcss -i cmd/web/assets/css/input.css -o cmd/web/assets/css/output.css

echo "Installation complete! You can now run ./start.sh to start the application."
