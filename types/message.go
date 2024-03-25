package types

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
	Map  string `json:"map"`
	Id   string `json:"id"`
}

