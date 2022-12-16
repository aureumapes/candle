# builds for linux
GOOS=linux go build -o build/candle-linux

# builds for win
GOOS=windows go build -o build/candle-windows.exe
