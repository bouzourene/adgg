mkdir -p ./build

env GOOS=windows GOARCH=amd64 go build -o build/adgg-win-amd64.exe
env GOOS=darwin GOARCH=amd64 go build -o build/adgg-osx-amd64
env GOOS=darwin GOARCH=arm64 go build -o build/adgg-osx-arm64
env GOOS=linux GOARCH=amd64 go build -o build/adgg-lin-amd64
env GOOS=linux GOARCH=arm64 go build -o build/adgg-lin-arm64

echo "Done"