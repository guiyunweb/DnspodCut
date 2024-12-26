package main

import (
	"DnspodCut/internal"
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	log.Println("项目开始...")

	// 读取配置文件
	config, err := internal.LoadYaml()
	if err != nil {
		log.Fatal("配置文件读取错误，程序停止......", err)
	}

	// 执行定时任务
	c := cron.New()
	err = c.AddFunc("0 * * * *", func() {
		log.Println("定时任务开始")
		internal.MonitoringAndUpdateDNS(config)
	})

	//开始定时任务
	c.Start()
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
