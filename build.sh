#!/bin/bash

echo "> Start building"

# Build Windows amd64
mkdir -p ./build/win-amd64
env GOOS=windows GOARCH=amd64 go build -o ./build/win-amd64/adgg.exe

# Build macOS Intel
mkdir -p ./build/osx-amd64
env GOOS=darwin GOARCH=amd64 go build -o ./build/osx-amd64/adgg
chmod +x ./build/osx-amd64/adgg

# Build macOS Apple Silicon
mkdir -p ./build/osx-arm64
env GOOS=darwin GOARCH=arm64 go build -o ./build/osx-arm64/adgg
chmod +x ./build/osx-arm64/adgg

# Build Linux amd64
mkdir -p ./build/linux-amd64
env GOOS=linux GOARCH=amd64 go build -o ./build/linux-amd64/adgg
chmod +x ./build/linux-amd64/adgg

# Build Linux arm64
mkdir -p ./build/linux-arm64
env GOOS=linux GOARCH=arm64 go build -o ./build/linux-arm64/adgg
chmod +x ./build/linux-arm64/adgg

echo "> Done building"
echo "> Start compression"

zip ./build/build-$(date +"%Y%m%d%H%M").zip ./build/*/*

echo "> Done with compression"

echo "> Cleaning up"
find ./build ! -name '*.zip' -type f -exec rm -f {} +
find ./build ! -name 'build' -type d -exec rmdir {} +