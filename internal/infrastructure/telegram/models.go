package telegram

func NewTGMessageInfo(Title TitleInfo, ChatID ChatID) TGMessageInfo {
	return TGMessageInfo{Title: Title, ID: ChatID}
}

type TGMessageInfo struct {
	Title TitleInfo
	ID    ChatID
}

type TitleInfo string
type ChatID int64

func (t TitleInfo) String() string {
	return string(t)
}

func (t ChatID) Int64() int64 {
	return int64(t)
}
