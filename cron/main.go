package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())
	c.Start()
	c.AddFunc("*/3 * * * * *", func() {
		fmt.Println("CRON")
	})
	select {}
}
