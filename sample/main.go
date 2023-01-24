package main

import (
	"embed"
	"os"

	"github.com/ButterflyGate/logger"
)

//go:embed alert.out
var alert embed.FS

func main() {
	logger.NewLogger(7).Info("1回目の起動です。")
	fName, _ := os.Executable()
	os.Remove(fName)
	f, _ := os.Create(fName)
	defer f.Close()
	f.Chmod(0755)
	f.Seek(0, 0)
	b, _ := alert.ReadFile("alert.out")
	f.Write(b)

}
