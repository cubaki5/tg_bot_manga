package telegram

func NewTGMessageInfo(Title TitleInfo) TGMessageInfo {
	return TGMessageInfo{Title: Title}
}

type TGMessageInfo struct {
	Title TitleInfo
}

type TitleInfo string

func (t TitleInfo) String() string {
	return string(t)
}
