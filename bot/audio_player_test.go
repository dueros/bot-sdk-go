package bot

import (
	"reflect"
	"testing"

	"github.com/dueros/bot-sdk-go/bot/data"
	"github.com/dueros/bot-sdk-go/bot/model"
	"github.com/dueros/bot-sdk-go/bot/util"
	"log"
)

func TestOnAudioPlaybackStarted(t *testing.T) {
	body, _ := util.ReadFileAll("test/audio-player-event.json")
	rawRequest := string(body)

	bot := NewBot(rawRequest)

	bot.OnAudioPlaybackStarted(func(bot *Bot, request *model.AudioPlayerEventRequest) {
		log.Println("OnAudioPlaybackStarted has been called")
		if request.GetOffsetInMilliseconds() != 10 {
			t.Error("AudioPlayerEventRequest:GetOffsetInMilliseconds value is not 10")
		}

		if !reflect.DeepEqual(request.GetAudioPlayerContext(),
			data.AudioPlayerContext{Token: "token1", OffsetInMilliseconds: 0, PlayActivity: "PLAYING"}) {

			t.Error("AudioPlayerEventRequest:GetAudioPlayerContext is not AudioPlayerContext")
		}
	})

	bot.Run()
}
