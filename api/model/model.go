package model

type ResTitle struct {
	Title string
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
