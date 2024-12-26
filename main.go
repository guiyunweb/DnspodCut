package main

import (
	"DnspodCut/utils"
)

func main() {
	//internal.LoadYaml()

	host := "124.113.13.12" // 你可以替换为任何活动的 IP 地址

	utils.Ping(host)

}
