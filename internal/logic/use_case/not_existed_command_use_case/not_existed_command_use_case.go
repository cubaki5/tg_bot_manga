package not_existed_command_use_case

type NotExistedCommandUseCase struct {
}

func NewNotExistedCommandUseCase() *NotExistedCommandUseCase {
	return &NotExistedCommandUseCase{}
}

func (s *NotExistedCommandUseCase) NotExistedCommand() string {
	startMessage :=
		`You should use the following commands:
				/start - to start working with me
				/add <title> - to add the new title to track
				/delete <title> - to delete the title from tracking list
				/list - to list all tracking titles`
	return startMessage
}
