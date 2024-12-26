package utils

import (
	"log"
	"net"
	"os/exec"
	"strings"
)

func Ping(host string) bool {

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Println("获取 IP地址错误:", err)
		return false
	}

	// 打印获取的 IP 地址
	for _, ip := range ips {
		log.Printf("域名 %s (IP: %s)\n", host, ip.String())
	}

	// 执行 ping 命令
	cmd := exec.Command("ping", "-c", "4", host) // 在 macOS 或 Linux 上使用 "-c" 参数，在 Windows 上使用 "-n" 参数
	stdout, err := cmd.Output()

	if err != nil {
		log.Println("Error executing ping command:", err)
	}

	output := string(stdout)
	lines := strings.Split(output, "\n")

	// 获取 ping 值
	var pingValue string
	for _, line := range lines {
		if strings.Contains(line, "time=") {
			pingValue = strings.Split(line, "time=")[1]
			break
		}
	}

	log.Printf("Ping: %s", lines)
	if pingValue != "" {
		log.Printf("Ping成功 %s. Ping值为: %s\n", host, pingValue)
		return true
	} else {
		log.Printf("Ping失败 %s .\n", host)
		return false
	}
}
