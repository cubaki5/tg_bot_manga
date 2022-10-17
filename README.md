## TG_BOT_MANGA

___

TG_BOT_MANGA is telegram bot for sending notifications to the user when the content of the tracking titles is updated.

## FUNCTIONALITY

___

At the moment, the TG_BOT_MANGA is single-user bot and works with only one manga web service. However, it can be easily expanded to work with other manga reading services. 

Currently, the TG_BOT_MANGA offers these commands:


- `/start` - tells the user how to use this bot</li>
- `/add` - adds a new title to the tracking list by link</li> 
- `/list` - returns a list of tracked titles</li>
- `/delete` - deletes the title from tracking list with the specified title</li>


After adding a title to the tracking list, the bot checks it for updates every hour. And if a new chapter is released, bot sends a notification to the user with the title name and a link to it.


## CODE INFORMATION

___

The TG_BOT_MANGA is implemented in the go programming language using a pure architecture design pattern. The project includes unit testing, where test coverage is 75.6%. 

The following external libraries are used in project:

- [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api/)
- [goquery](https://github.com/PuerkitoBio/goquery)
- [gomock](https://github.com/golang/mock)
- [testify](https://github.com/stretchr/testify)

The golangci-lint linters involved in the project are: goimports, goconst, gosec, govet, ineffassign, revive, typecheck, unused. Also a pre-commit hook for golangci-lint is included in project.