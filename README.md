## TG_BOT_MANGA

<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="96" height="20" role="img" aria-label="coverage: 76%">
<title>coverage: 76%</title>
<linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><clipPath id="r"><rect width="96" height="20" rx="3" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="61" height="20" fill="#555"/><rect x="61" width="35" height="20" fill="#4c1"/><rect width="96" height="20" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110"><text aria-hidden="true" x="315" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="510">coverage</text>
<text x="315" y="140" transform="scale(.1)" fill="#fff" textLength="510">coverage</text>
<text aria-hidden="true" x="775" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="250">76%</text>
<text x="775" y="140" transform="scale(.1)" fill="#fff" textLength="250">76%</text></g>
</svg>

TG_BOT_MANGA is telegram bot for sending notifications to the user when the content of the tracking titles is updated.

## FUNCTIONALITY


At the moment, the TG_BOT_MANGA is the single-user bot and works with only one manga web service. However, it can be easily expanded to be multi-users and to work with other manga reading services. 

Currently, the TG_BOT_MANGA offers these commands:


- `/start` - tells the user how to use this bot</li>
- `/add` - adds a new title to the tracking list by link</li> 
- `/list` - returns a list of tracked titles</li>
- `/delete` - deletes the title from tracking list with the specified title</li>


After adding a title to the tracking list, the bot checks it for updates every hour. And if a new chapter is released, bot sends a notification to the user with the title name and a link to it.


## CODE INFORMATION

The TG_BOT_MANGA is implemented in the go programming language using a pure architecture design pattern. The project includes unit testing, where test coverage is 75.6%. 

The following external libraries are used in project:

- [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api/)
- [goquery](https://github.com/PuerkitoBio/goquery)
- [gomock](https://github.com/golang/mock)
- [testify](https://github.com/stretchr/testify)

The golangci-lint linters involved in the project are: goimports, goconst, gosec, govet, ineffassign, revive, typecheck, unused.

A pre-commit hook includes run of golangci-lint and unit tests.