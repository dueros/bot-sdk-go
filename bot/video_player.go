package bot

import (
	"github.com/dueros/bot-sdk-go/bot/model"
)

// 视频播放相关的事件处理注册
// 设备在播放视频事会上报event，具体可以参考DCS.video_player.play事件，TODO
// 技能收到的事件数据格式参考技能协议，TODO

// 视频开始播放事件
func (this *Bot) OnVideoPlaybackStarted(fn func(bot *Bot, request *model.VideoPlayerEventRequest)) {
	this.AddEventListener(model.VIDEO_PLAYER_PLAYBACK_STARTED, func(bot *Bot, request interface{}) {
		req := request.(model.VideoPlayerEventRequest)
		fn(bot, &req)
	})
}

// 视频停止播放事件
func (this *Bot) OnVideoPlaybackStopped(fn func(bot *Bot, request *model.VideoPlayerEventRequest)) {
	this.AddEventListener(model.VIDEO_PLAYER_PLAYBACK_STOPPED, func(bot *Bot, request interface{}) {
		req := request.(model.VideoPlayerEventRequest)
		fn(bot, &req)
	})
}

// 视频播放完成事件
func (this *Bot) OnVideoPlaybackFinished(fn func(bot *Bot, request *model.VideoPlayerEventRequest)) {
	this.AddEventListener(model.VIDEO_PLAYER_PLAYBACK_FINISHED, func(bot *Bot, request interface{}) {
		req := request.(model.VideoPlayerEventRequest)
		fn(bot, &req)
	})
}

// 视频快要播放结束上报的事件
func (this *Bot) OnVideoPlaybackNearlyFinished(fn func(bot *Bot, request *model.VideoPlayerEventRequest)) {
	this.AddEventListener(model.VIDEO_PLAYER_PLAYBACK_NEARLY_FINISHED, func(bot *Bot, request interface{}) {
		req := request.(model.VideoPlayerEventRequest)
		fn(bot, &req)
	})
}

// 视频周期上报播放进度
func (this *Bot) OnVideoRrogressReportIntevalElapsed(fn func(bot *Bot, request *model.VideoPlayerEventRequest)) {
	this.AddEventListener(model.VIDEO_PLAYER_PROGRESS_REPORT_INTERVAL_ELAPSED, func(bot *Bot, request interface{}) {
		req := request.(model.VideoPlayerEventRequest)
		fn(bot, &req)
	})
}

// 视频自动暂停后上报
func (this *Bot) OnVideoPlayerScheduledStopReached(fn func(bot *Bot, request *model.VideoPlayerEventRequest)) {
	this.AddEventListener(model.VIDEO_PLAYER_PLAYBACK_SCHEDULED_STOP_REACHED, func(bot *Bot, request interface{}) {
		req := request.(model.VideoPlayerEventRequest)
		fn(bot, &req)
	})
}
