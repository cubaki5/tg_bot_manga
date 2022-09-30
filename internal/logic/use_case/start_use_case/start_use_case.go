package start_use_case

type StartUseCase struct {
}

func NewStartUseCase() *StartUseCase {
	return &StartUseCase{}
}

func (s *StartUseCase) Start() string {
	startMessage :=
		`Hello, I'm telegram bot for notifying about manga updates.
		I have the following commands:
		/start - to start working with me
		/add <title url> - to add the new title to track
		/delete <title> - to delete the title from tracking list
		/list - to list all tracking titles`
	return startMessage
}
