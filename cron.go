package main

import (
	"github.com/robfig/cron"
	"log"
)

func main() {
	i := 0
	c := cron.New()
	s := "0 */1 * * * *"
	c.AddFunc(s, func() {
		i++
		log.Println("number ", i)
	})
	c.Start()
	select {}
}
