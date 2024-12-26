package internal

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	SecretId  string `yaml:"secretId"`
	SecretKey string `yaml:"secretKey"`
	Dns       Dns    `yaml:"dns"`
}

type Dns struct {
	Domain       string `yaml:"domain"`
	RecordType   string `yaml:"recordType"`
	RecordLine   string `yaml:"recordLine"`
	Value        string `yaml:"value"`
	SubDomain    string `yaml:"subdomain"`
	RecordLineId string `yaml:"recordLineId"`
}

func LoadYaml() {
	dataBytes, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("yaml 文件的内容: \n", string(dataBytes))
	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	fmt.Printf("config → %+v\n", config) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}

	mp := make(map[string]any, 2)
	err = yaml.Unmarshal(dataBytes, mp)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	fmt.Printf("map → %+v", config) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}

}
