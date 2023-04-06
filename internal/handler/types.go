package handler

type segmentInfo struct {
	LangKind string  `json:"langKind"`
	AsrKind  string  `json:"asrKind"`
	Id       string  `json:"id"`
	Start    float64 `json:"start"`
	Length   float64 `json:"length"`
}

type youtubeInfo struct {
	AsrKind string `json:"asrKind"`
	Vid     string `json:"vid"`
}

type result struct {
	Done bool `json:"done"`
	Data any  `json:"data"`
}
