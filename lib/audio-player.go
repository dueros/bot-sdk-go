package lib

import (
	"github.com/dueros/bot-sdk-go/lib/req"
)

// 音频播放相关的事件处理注册
// 设备在播放音频事会上报event，具体可以参考DCS.audio.play事件，TODO
// 技能收到的事件数据格式参考技能协议，TODO

// 音频开始播放事件
func (this *Bot) OnPlaybackStarted(fn func(bot *Bot, request *req.EventRequest)) {
	this.AddEventListener(req.AUDIO_PLAYER_PLAYBACK_STARTED, fn)
}

// 音频停止播放事件
func (this *Bot) OnPlaybackStopped(fn func(bot *Bot, request *req.EventRequest)) {
	this.AddEventListener(req.AUDIO_PLAYER_PLAYBACK_STARTED, fn)
}

// 音频播放完成事件
func (this *Bot) OnPlaybackFinished(fn func(bot *Bot, request *req.EventRequest)) {
	this.AddEventListener(req.AUDIO_PLAYER_PLAYBACK_FINISHED, fn)
}

// 音频快要播放结束上报的事件
func (this *Bot) OnPlaybackNearlyFinished(fn func(bot *Bot, request *req.EventRequest)) {
	this.AddEventListener(req.AUDIO_PLAYER_PLAYBACK_NEARLY_FINISHED, fn)
}

func (this *Bot) OnRrogressReportIntevalElapsed(fn func(bot *Bot, request *req.EventRequest)) {
	this.AddEventListener(req.AUDIO_PLAYER_PROGRESS_REPORT_INTERVAL_ELAPSED, fn)
}
