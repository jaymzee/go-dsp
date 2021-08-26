package main

import "os"

func isatty() bool {
	fileInfo, _ := os.Stdout.Stat()
	return fileInfo.Mode()&os.ModeCharDevice != 0
}
