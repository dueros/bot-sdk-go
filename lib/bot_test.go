package lib

import "testing"

func TestAddIntentHandler(t *testing.T) {
	bot := NewBot()
	bot.AddIntentHandler("example", func() ResponseData {
		return ResponseData{Card: "test"}
	})
	_, ok := bot.intentHandler["example"]
	if !ok {
		t.Errorf("AddIntentHandler Error")
	}
}
