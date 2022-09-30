package start_use_case

import "fmt"

type StartUseCase struct {
}

func NewStartUseCase() *StartUseCase {
	return &StartUseCase{}
}

func (s *StartUseCase) Start() string {
	startMessage := fmt.Sprintf("Hello, I'm telegram bot for notifying about manga updates.\nI have the following commands:\n/start - to start working with me\n/add <title> - to add the new title to track\n/delete <title> - to delete the title from tracking list\n/list - to list all tracking titles")
	return startMessage
}
