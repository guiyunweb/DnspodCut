package utils

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Ping(ipAddress string) {
	// 要 ping 的 IP 地址

	// 执行 ping 命令
	cmd := exec.Command("ping", "-c", "4", ipAddress) // 在 macOS 或 Linux 上使用 "-c" 参数，在 Windows 上使用 "-n" 参数
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

	if pingValue != "" {
		log.Printf("Ping to %s is successful. Ping value: %s\n", ipAddress, pingValue)
		fmt.Printf("Ping to %s is successful. Ping value: %s\n", ipAddress, pingValue)
	} else {
		log.Printf("Ping to %s failed\n", ipAddress)
		fmt.Printf("Ping to %s failed\n", ipAddress)
	}
}
