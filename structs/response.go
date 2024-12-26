package structs

type ResponseSelectDns struct {
	RecordList []RecordList `json:"RecordList"`
	Age        int          `json:"age"`
}

type RecordList struct {
	RecordId uint64 `json:"RecordId"`
	Type     string `json:"Type"`
	Value    string `json:"Value"`
	Status   string `json:"Status"`
	Line     string `json:"Line"`
}
