package main

import (
	dueros "github.com/dueros/bot-sdk-go/lib"
	card "github.com/dueros/bot-sdk-go/lib/card"
	audio "github.com/dueros/bot-sdk-go/lib/directive/audio"
	req "github.com/dueros/bot-sdk-go/lib/req"
)

func main() {

	application := dueros.Application{Handler: func(rawData string) string {
		bot := dueros.NewBot(rawData)
		bot.AddIntentHandler("test.inquiry888",
			func(this *dueros.Bot, request *req.IntentRequest) {
				card := card.NewStandardCard()
				card.SetTitle("gasg")
				card.SetImage("gasg")
				request.Dialog.ConfirmIntent()
				this.Response.Command(audio.NewPlayDirective()).Ask("gagagagag").DisplayCard(card)
			})

		return bot.Run()
	}}
	application.Start("localhost:8888")
}
