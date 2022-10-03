package models_types

type (
	TitleName string
	URL       string
	TitleID   int64
)

func (t TitleName) String() string {
	return string(t)
}

func (u URL) String() string {
	return string(u)
}

func (t TitleID) Int64() int64 {
	return int64(t)
}
