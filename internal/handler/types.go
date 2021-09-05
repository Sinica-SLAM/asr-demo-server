package handler

type segmentInfo struct {
	LangKind string  `json:"langKind"`
	AsrKind  string  `json:"asrKind"`
	Id       string  `json:"id"`
	Start    float64 `json:"start"`
	Length   float64 `json:"length"`
}

type uploadInfo struct {
	LangKind string `form:"langKind"`
	AsrKind  string `form:"asrKind"`
	File     string `form:"file" format:"binary"`
}

type wordalignment struct {
	Start  float64 `json:"start"`
	Length float64 `json:"length"`
	Word   string  `json:"word"`
	Token  string  `json:"token"`
}
