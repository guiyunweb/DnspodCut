package structs

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
