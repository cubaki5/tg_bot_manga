package start_handler

type UseCase interface {
	Start() error
}

type StartHandler struct {
	uc UseCase
}

func NewStartHandler(uc UseCase) *StartHandler {
	return &StartHandler{uc: uc}
}

func (a StartHandler) Handle() error {
	return nil
}
