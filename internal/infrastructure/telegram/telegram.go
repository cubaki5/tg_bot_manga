package telegram

type Handler interface {
	Handle() error
}

type TelegramBot struct {
	handlers map[string]Handler
}

func NewTGBot(handlers map[string]Handler) *TelegramBot {
	return &TelegramBot{
		handlers: handlers,
	}
}

func (t *TelegramBot) Run() {
	msg := "/start"
	t.handlers[msg].Handle()
}
