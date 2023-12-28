package model

type TitleInfo struct {
	Titles []string
}

type TransMode struct {
	ModeID int
}

type Receive struct {
	ApiID  string
	ModeID int // setMode の 遷移先を設定するときだけ
}

type Result struct {
	Status string
}

func NewTitle() *TitleInfo {
	return &TitleInfo{}
}

func NewTitleHandler() *TitleHandler {
	return &TitleHandler{TitleInfo: &TitleInfo{}}
}

type TitleHandler struct {
	*TitleInfo
}
