package main

import (
	dueros "github.com/dueros/bot-sdk-go/bot"
	"github.com/dueros/bot-sdk-go/bot/card"
	"github.com/dueros/bot-sdk-go/bot/model"
)

func main() {

	application := dueros.Application{Handler: func(body string) string {
		bot := dueros.NewBot(body)
		bot.AddIntentHandler("test.inquiry888",
			func(this *dueros.Bot, request *model.IntentRequest) {
				c := card.NewTextCard("欢迎查询")

				request.Dialog.ConfirmIntent()

				this.Response.Ask("欢迎查询").DisplayCard(c)
			})

		return bot.Run()
	}}
	application.Start(":8888")
}
