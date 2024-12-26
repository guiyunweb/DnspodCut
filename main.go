package main

import (
	"DnspodCut/internal"
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	log.Println("项目开始...")

	internal.LoadYaml()

	c := cron.New()

	err := c.AddFunc("* * * * *", func() {
		log.Println("定时任务执行")
	})
	if err != nil {
		return
	}

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
