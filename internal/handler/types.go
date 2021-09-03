package handler

type segmentInfo struct {
	LangKind string  `json:"langKind"`
	AsrKind  string  `json:"asrKind"`
	Id       string  `json:"id"`
	Start    float64 `json:"start"`
	Length   float64 `json:"length"`
}
