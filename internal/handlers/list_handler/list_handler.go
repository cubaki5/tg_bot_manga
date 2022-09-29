package list_handler

type UseCase interface {
	List() (interface{}, error)
}

type ListHandler struct {
	uc UseCase
}

func NewListHandler(uc UseCase) *ListHandler {
	return &ListHandler{uc: uc}
}

func (gh *ListHandler) Handle() error {
	return nil
}
