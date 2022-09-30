package delete_handler

type UseCase interface {
	Delete() error
}

type DeleteHandler struct {
	uc UseCase
}

func NewDeleteHandler(uc UseCase) *DeleteHandler {
	return &DeleteHandler{uc: uc}
}

func (dh *DeleteHandler) Handle() error {
	return nil
}
