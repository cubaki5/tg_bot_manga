package telegram

func NewTGMessageInfo(title TitleInfo) TGMessageInfo {
	return TGMessageInfo{title: title}
}

type TGMessageInfo struct {
	title TitleInfo
}

type TitleInfo string
