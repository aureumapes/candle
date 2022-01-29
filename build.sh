# builds for windows
env GOOS=linux GOARCH=amd64 go build -o build/candle

# builds for linux
env GOOS=windows GOARCH=amd64 go build -o build/candle.exe