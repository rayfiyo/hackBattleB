package utils

import (
	"api/model"
)

// とりあえず，レスポンスだけ実装

func ReadTitle() model.TitleInfo {
	title := model.TitleInfo{
		Title: "バンジージャンプ",
	}
	return title
}

func ReadMode() model.TransMode {
	mode := model.TransMode{
		ModeID: 0,
	}
	return mode
}

func WriteMode() model.Result {
	status := model.Result{
		Status: "OK",
	}
	return status
}

func Clear() model.Result {
	status := model.Result{
		Status: "OK",
	}
	return status
}
