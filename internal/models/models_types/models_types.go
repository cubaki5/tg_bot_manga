package models_types

type (
	TitleName string
	URL       string
	TitleID   uint32

	ChatID int64
)

func (t TitleName) String() string {
	return string(t)
}

func (u URL) String() string {
	return string(u)
}

func (t TitleID) Uint32() uint32 {
	return uint32(t)
}

func (t ChatID) Int64() int64 {
	return int64(t)
}
