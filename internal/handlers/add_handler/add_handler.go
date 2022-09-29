package add_handler

type UseCase interface {
	Add() error
}

type AddHandler struct {
	uc UseCase
}

func NewAddHandler(uc UseCase) *AddHandler {
	return &AddHandler{uc: uc}
}

func (a AddHandler) Handle() error {
	//TODO implement me
	panic("implement me")
}
