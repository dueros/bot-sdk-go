package bot

import (
	//"reflect"
	"testing"

	"github.com/dueros/bot-sdk-go/bot/model"
	"github.com/dueros/bot-sdk-go/bot/util"
	"log"
)

func TestOnVideoPlaybackStarted(t *testing.T) {
	body, _ := util.ReadFileAll("test/video-player-event.json")
	rawRequest := string(body)

	bot := NewBot(rawRequest)

	bot.OnVideoPlaybackStarted(func(bot *Bot, request *model.VideoPlayerEventRequest) {
		log.Println("OnVideoPlaybackStarted has been called")
		if request.GetOffsetInMilliseconds() != 10 {
			t.Error("VideoPlayerEventRequest:GetOffsetInMilliseconds value is not 10")
		}
	})

	bot.Run()
}
