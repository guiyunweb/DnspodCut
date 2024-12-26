package internal

import (
	"DnspodCut/structs"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func LoadYaml() (structs.Config, error) {
	dataBytes, err := os.ReadFile("config.yaml")
	config := structs.Config{}
	if err != nil {
		log.Println("读取文件失败：", err)
		return config, err
	}

	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		log.Println("解析 yaml 文件失败：", err)
		return config, err
	}
	log.Printf("config → %+v\n", config)

	mp := make(map[string]any, 2)
	err = yaml.Unmarshal(dataBytes, mp)
	if err != nil {
		log.Println("解析 yaml 文件失败：", err)
		return config, err
	}

	log.Printf("配置读取成功")
	return config, nil
}
