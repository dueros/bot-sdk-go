package lib

import (
	"github.com/dueros/bot-sdk-go/lib/req"
)

func (this *Bot) OnPlaybackStarted(fn func(bot *Bot, request *req.EventRequest)) {
	this.AddEventListener(req.AUDIO_PLAYER_PLAYBACK_STARTED, fn)
}

func (this *Bot) OnPlaybackStopped(fn func(bot *Bot, request *req.EventRequest)) {
	this.AddEventListener(req.AUDIO_PLAYER_PLAYBACK_STARTED, fn)
}

func (this *Bot) OnPlaybackFinished(fn func(bot *Bot, request *req.EventRequest)) {
	this.AddEventListener(req.AUDIO_PLAYER_PLAYBACK_FINISHED, fn)
}

func (this *Bot) OnPlaybackNearlyFinished(fn func(bot *Bot, request *req.EventRequest)) {
	this.AddEventListener(req.AUDIO_PLAYER_PLAYBACK_NEARLY_FINISHED, fn)
}

func (this *Bot) OnRrogressReportIntevalElapsed(fn func(bot *Bot, request *req.EventRequest)) {
	this.AddEventListener(req.AUDIO_PLAYER_PROGRESS_REPORT_INTERVAL_ELAPSED, fn)
}
