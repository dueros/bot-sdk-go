package lib

import "testing"

func TestAddIntentHandler(t *testing.T) {
	bot := NewBot()
	bot.AddIntentHandler("example", func(bot *Bot, request *req.IntentRequest) {
		return
	})
	_, ok := bot.intentHandler["example"]
	if !ok {
		t.Errorf("AddIntentHandler Error")
	}
}
