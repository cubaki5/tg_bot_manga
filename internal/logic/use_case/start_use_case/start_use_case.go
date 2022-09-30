package start_use_case

type StartUseCase struct {
}

func NewStartUseCase() *StartUseCase {
	return &StartUseCase{}
}

func (s *StartUseCase) Start() error {
	return nil
}
