package internal

import (
	"DnspodCut/utils"
	"log"
)

var m = make(map[string]int)

func MonitoringAndUpdateDNS(config Config) {
	dns := config.Dns

	for _, item := range dns {

		domain := item.SubDomain + "." + item.Domain
		log.Printf("\n\n================= %s =================", domain)

		//查看DNS是否成功
		isPing := utils.Ping(domain)

		if !isPing {
			// 判断键是否存在
			if value, exists := m[domain]; exists {
				log.Printf("'%s' 不通，错误数值为 %d\n", domain, value)
				m[domain] = value + 1
			} else {
				log.Printf("'%s' 不通，错误数值为 %d\n", domain, value)
				m[domain] = 1
			}
			if m[domain] < 0 {
				log.Printf("'%s' 值切换中......\n", domain)
			}
		} else {
			m[domain] = 0
		}

		if m[domain] > config.ErrorNum {
			log.Printf("'%s'错误数 %d已超过,现为%d.开始切换  \n", domain, config.ErrorNum, m[domain])
			for _, value := range item.Value {
				isValue := utils.Ping(value)
				if isValue {
					utils.UpdateDns(config, item, domain)
					m[domain] = -30
				}
			}
		}

	}

}
