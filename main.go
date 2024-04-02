package main

import (
    "fmt"
	"os"

	"github.com/mymmrac/telego"

	th "github.com/mymmrac/telego/telegohandler"

	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := "bot"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err!= nil {
        fmt.Println(err)
        os.Exit(1)
    }

	updates, _:= bot.UpdatesViaLongPolling(nil)
	bh, _:= th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)

		keyboard := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("Number").WithRequestContact(),
				tu.KeyboardButton("Geo").WithRequestLocation(),
				tu.KeyboardButton("help").WithRequestUsers(),
			),
		)
		message := tu.Message(
			chatID,
			"Hello sexy girl!",
		).WithReplyMarkup(keyboard)
		
		_, _ = bot.SendSticker(
			tu.Sticker(
				chatID,
				tu.FileFromID("CAACAgQAAxkBAAEL0ilmCVgbflEOfe-Mgb830BWmTBSHiQACCQwAAktAyFBpFgYK5U7XWTQE"),
			),
		)
		_, _ = bot.SendMessage(message)
	}, th.CommandEqual("start"))
	
	bh.Start()
}