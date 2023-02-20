package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/schollz/progressbar/v3"
	"time"
)

var bar *progressbar.ProgressBar

func main() {
	bar = progressbar.NewOptions(1000,
		//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[blue][1/3][reset] Writing moshable file..."))
	//progressbar.OptionSetTheme(progressbar.Theme{
	//	Saucer:        "[green]=[reset]",
	//	SaucerHead:    "[green]>[reset]",
	//	SaucerPadding: " ",
	//	BarStart:      "[",
	//	BarEnd:        "]",
	//}))
	for i := 0; i < 1000; i++ {
		bar.Add(1)
		time.Sleep(5 * time.Millisecond)
	}
}
