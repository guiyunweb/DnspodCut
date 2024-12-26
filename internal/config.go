package internal

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	SecretId  string `yaml:"secretId"`
	SecretKey string `yaml:"secretKey"`
	ErrorNum  int    `yaml:"errorNum"`
	Dns       []Dns  `yaml:"dns"`
}

type Dns struct {
	Domain       string   `yaml:"domain"`
	RecordType   string   `yaml:"recordType"`
	RecordLine   string   `yaml:"recordLine"`
	Value        []string `yaml:"value"`
	SubDomain    string   `yaml:"subdomain"`
	RecordLineId string   `yaml:"recordLineId"`
}

func LoadYaml() (Config, error) {
	dataBytes, err := os.ReadFile("config.yaml")
	config := Config{}
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
