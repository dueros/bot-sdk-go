package bot

import (
	"github.com/dueros/bot-sdk-go/bot/util"
	"testing"
)

func TestAddDefaultEventListener(t *testing.T) {
	body, _ := util.ReadFileAll("test/audio-player-event.json")
	rawRequest := string(body)

	b := NewBot(rawRequest)
	called := false

	b.AddDefaultEventListener(func(bot *Bot, request interface{}) {
		called = true
		bot.Response.HoldOn()
		ret := bot.Response.GetData()
		shouldEndSession := ret["shouldEndSession"].(bool)

		if shouldEndSession != false {
			t.Error("AddDefaultEventListener HoldOn: shouldEndSession is not false")
		}
	})

	b.Run()

	if !called {
		t.Error("AddDefaultEventListener has not been called")
	}
}
