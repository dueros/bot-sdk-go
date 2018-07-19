package bot

import (
	"github.com/dueros/bot-sdk-go/bot/model"
)

// 音频播放相关的事件处理注册
// 设备在播放音频事会上报event，具体可以参考DCS.audio_player.play事件，TODO
// 技能收到的事件数据格式参考技能协议，TODO

// 音频开始播放事件
func (this *Bot) OnAudioPlaybackStarted(fn func(bot *Bot, request *model.AudioPlayerEventRequest)) {
	this.AddEventListener(model.AUDIO_PLAYER_PLAYBACK_STARTED, func(bot *Bot, request interface{}) {
		req := request.(model.AudioPlayerEventRequest)
		fn(bot, &req)
	})
}

// 音频停止播放事件
func (this *Bot) OnAudioPlaybackStopped(fn func(bot *Bot, request *model.AudioPlayerEventRequest)) {
	this.AddEventListener(model.AUDIO_PLAYER_PLAYBACK_STOPPED, func(bot *Bot, request interface{}) {
		req := request.(model.AudioPlayerEventRequest)
		fn(bot, &req)
	})
}

// 音频播放完成事件
func (this *Bot) OnAudioPlaybackFinished(fn func(bot *Bot, request *model.AudioPlayerEventRequest)) {
	this.AddEventListener(model.AUDIO_PLAYER_PLAYBACK_FINISHED, func(bot *Bot, request interface{}) {
		req := request.(model.AudioPlayerEventRequest)
		fn(bot, &req)
	})
}

// 音频快要播放结束上报的事件
func (this *Bot) OnAudioPlaybackNearlyFinished(fn func(bot *Bot, request *model.AudioPlayerEventRequest)) {
	this.AddEventListener(model.AUDIO_PLAYER_PLAYBACK_NEARLY_FINISHED, func(bot *Bot, request interface{}) {
		req := request.(model.AudioPlayerEventRequest)
		fn(bot, &req)
	})
}

// 音频周期上报播放进度
func (this *Bot) OnAudioRrogressReportIntevalElapsed(fn func(bot *Bot, request *model.AudioPlayerEventRequest)) {
	this.AddEventListener(model.AUDIO_PLAYER_PROGRESS_REPORT_INTERVAL_ELAPSED, func(bot *Bot, request interface{}) {
		req := request.(model.AudioPlayerEventRequest)
		fn(bot, &req)
	})
}
