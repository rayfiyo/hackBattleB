package model

type ResTitle struct {
	Title string `json:"title"`
}

type TransMode struct {
	ModeID int `json:"mode_id"`
}

type Receive struct {
	ApiID  string `json:"api_id"`
	ModeID int    `json:"mode_id"` // setMode の 遷移先を設定するときだけ
}

type Result struct {
	Status string `json:"status"`
}
