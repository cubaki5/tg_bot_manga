package main

import (
	"log"

	"tgbot/internal/handlers/add_handler"
	"tgbot/internal/handlers/delete_handler"
	"tgbot/internal/handlers/list_handler"
	"tgbot/internal/handlers/not_existed_command_handler"
	"tgbot/internal/handlers/start_handler"
	"tgbot/internal/infrastructure/mint_client"
	"tgbot/internal/infrastructure/runtime_database"
	"tgbot/internal/infrastructure/telegram"
	"tgbot/internal/logic/modules/mint_information"
	"tgbot/internal/logic/modules/mint_information/parsers/HTML"
	"tgbot/internal/logic/modules/notifier"
	"tgbot/internal/logic/use_case/add_use_case"
	"tgbot/internal/logic/use_case/delete_use_case"
	"tgbot/internal/logic/use_case/list_use_case"
	"tgbot/internal/logic/use_case/not_existed_command_use_case"
	"tgbot/internal/logic/use_case/start_use_case"
)

func main() {
	db := runtime_database.NewDatabase()
	mintClient := mint_client.NewMintClient()
	parser := HTML.NewParser()
	tgBot, err := telegram.NewTGBot()
	if err != nil {
		log.Fatal(err)
	}

	not := notifier.NewNotifier(mintClient, db, tgBot, parser)
	getModule := mint_information.NewGetTitleModule(mintClient, parser)

	adUC := add_use_case.NewAddUseCase(db, getModule)
	delUC := delete_use_case.NewDeleteUseCase(db)
	listUC := list_use_case.NewListUseCase(db)
	startUC := start_use_case.NewStartUseCase()
	notExecComUC := not_existed_command_use_case.NewNotExistedCommandUseCase()

	tgBot.AppendBotWithHandlers(map[string]telegram.Handler{
		"start":  start_handler.NewStartHandler(startUC),
		"add":    add_handler.NewAddHandler(adUC),
		"delete": delete_handler.NewDeleteHandler(delUC),
		"list":   list_handler.NewListHandler(listUC),
	}, not_existed_command_handler.NewNotExistedCommandHandler(notExecComUC))

	tgBot.Run()
	not.CheckUpdates()
}
