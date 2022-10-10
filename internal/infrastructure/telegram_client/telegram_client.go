package telegram_client

import "tgbot/internal/models"

type TelegramClient struct{}

func NewTelegramClient() *TelegramClient {
	return &TelegramClient{}
}

func (t *TelegramClient) PostMsg(title models.Title) error {
	//TODO implement me
	panic("implement me")
}
