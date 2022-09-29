package main

import (
	"tgbot/internal/handlers/add_handler"
	"tgbot/internal/handlers/delete_handler"
	"tgbot/internal/handlers/list_handler"
	"tgbot/internal/handlers/start_handler"
	"tgbot/internal/infrastructure/mint_client"
	"tgbot/internal/infrastructure/runtime_database"
	"tgbot/internal/infrastructure/telegram"
	"tgbot/internal/infrastructure/telegram_client"
	"tgbot/internal/logic/modules/notifier"
	"tgbot/internal/logic/use_case/add_use_case"
	"tgbot/internal/logic/use_case/delete_use_case"
	"tgbot/internal/logic/use_case/list_use_case"
	"tgbot/internal/logic/use_case/start_use_case"
)

func main() {
	db := runtime_database.NewDatabase(map[string]interface{}{})
	webClient := mint_client.NewMintClient()
	tgClient := telegram_client.NewTelegramClient()

	not := notifier.NewNotifier(webClient, db, tgClient)

	adUC := add_use_case.NewAddUseCase(db)
	delUC := delete_use_case.NewDeleteUseCase(db)
	getUC := list_use_case.NewGetUseCase(db)
	startUC := start_use_case.NewStartUseCase()

	tgBot := telegram.NewTGBot(map[string]telegram.Handler{
		"/start":  start_handler.NewStartHandler(startUC),
		"/add":    add_handler.NewAddHandler(adUC),
		"/delete": delete_handler.NewDeleteHandler(delUC),
		"/get":    list_handler.NewListHandler(getUC),
	})

	tgBot.Run()
	not.CheckUpdates()
	//кого-то точно запускаю в отдельной горутине
}
