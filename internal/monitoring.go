package internal

import (
	"DnspodCut/structs"
	"DnspodCut/utils"
	"log"
	"strconv"
)

var m = make(map[string]int)

func MonitoringAndUpdateDNS(config structs.Config) {
	dns := config.Dns

	for _, item := range dns {

		domain := item.SubDomain + "." + item.Domain

		log.Printf("\n\n================= %s =================", domain)

		data, err := utils.FindDns(config, item)
		if err != nil {
			log.Printf("请求DNS列表失败 \n")
		}

		for _, record := range data {
			isPing := utils.Ping(record.Value)
			if !isPing {
				if value, exists := m[strconv.FormatUint(record.RecordId, 10)]; exists {
					log.Printf("'%s' 不通，错误数值为 %d\n", record.Value, value)
					m[strconv.FormatUint(record.RecordId, 10)] = value + 1
				} else {
					log.Printf("'%s' 不通，错误数值为 %d\n", record.Value, value)
					m[strconv.FormatUint(record.RecordId, 10)] = 1
				}

			} else {
				if record.Status == "DISABLE" {
					utils.UpdateDns(config, item, record, "ENABLE")
					log.Printf("'%s' 恢复正常，将解析修改为开启\n", record.Value)
				} else {
					log.Printf("'%s' 正常......\n", record.Value)

				}
			}

			if m[strconv.FormatUint(record.RecordId, 10)] > config.ErrorNum {
				utils.UpdateDns(config, item, record, "DISABLE")
				log.Printf("'%s' 请求多次异常，将解析修改为暂停\n", record.Value)
			}
		}

	}

}
