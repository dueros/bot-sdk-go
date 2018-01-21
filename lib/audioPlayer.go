package lib

import (
	"./req"
)

func (this *Bot) OnPlaybackStarted(fn func(bot *Bot, request *req.EventRequest)) {
	this.AddEventListener(req.AUDIO_PLAYER_PLAYBACK_STARTED, fn)
}
