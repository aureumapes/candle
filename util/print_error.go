package util

import "os"

func PrintError(error string) {
	println("\033[31mAn error occured with Candle:\x1b[91m\n" + error + "\033[0m")
	os.Exit(0)
}
