

# bot
`import "github.com/dueros/bot-sdk-go/bot"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type Bot](#Bot)
  * [func NewBot(rawRequest string) *Bot](#NewBot)
  * [func (this *Bot) AddDefaultEventListener(fn func(bot *Bot, request interface{}))](#Bot.AddDefaultEventListener)
  * [func (this *Bot) AddEventListener(eventName string, fn func(bot *Bot, request interface{}))](#Bot.AddEventListener)
  * [func (this *Bot) AddIntentHandler(intentName string, fn func(bot *Bot, request *model.IntentRequest))](#Bot.AddIntentHandler)
  * [func (this *Bot) OnAudioPlaybackFinished(fn func(bot *Bot, request *model.AudioPlayerEventRequest))](#Bot.OnAudioPlaybackFinished)
  * [func (this *Bot) OnAudioPlaybackNearlyFinished(fn func(bot *Bot, request *model.AudioPlayerEventRequest))](#Bot.OnAudioPlaybackNearlyFinished)
  * [func (this *Bot) OnAudioPlaybackStarted(fn func(bot *Bot, request *model.AudioPlayerEventRequest))](#Bot.OnAudioPlaybackStarted)
  * [func (this *Bot) OnAudioPlaybackStopped(fn func(bot *Bot, request *model.AudioPlayerEventRequest))](#Bot.OnAudioPlaybackStopped)
  * [func (this *Bot) OnAudioRrogressReportIntevalElapsed(fn func(bot *Bot, request *model.AudioPlayerEventRequest))](#Bot.OnAudioRrogressReportIntevalElapsed)
  * [func (this *Bot) OnDisplayElementSelected(fn func(bot *Bot, request *model.EventRequest))](#Bot.OnDisplayElementSelected)
  * [func (this *Bot) OnLaunchRequest(fn func(bot *Bot, request *model.LaunchRequest))](#Bot.OnLaunchRequest)
  * [func (this *Bot) OnLinkAccountSuccessed(fn func(bot *Bot, request *model.EventRequest))](#Bot.OnLinkAccountSuccessed)
  * [func (this *Bot) OnScreenLinkClicked(fn func(bot *Bot, request *model.EventRequest))](#Bot.OnScreenLinkClicked)
  * [func (this *Bot) OnSessionEndedRequest(fn func(bot *Bot, request *model.SessionEndedRequest))](#Bot.OnSessionEndedRequest)
  * [func (this *Bot) OnVideoPlaybackFinished(fn func(bot *Bot, request *model.VideoPlayerEventRequest))](#Bot.OnVideoPlaybackFinished)
  * [func (this *Bot) OnVideoPlaybackNearlyFinished(fn func(bot *Bot, request *model.VideoPlayerEventRequest))](#Bot.OnVideoPlaybackNearlyFinished)
  * [func (this *Bot) OnVideoPlaybackStarted(fn func(bot *Bot, request *model.VideoPlayerEventRequest))](#Bot.OnVideoPlaybackStarted)
  * [func (this *Bot) OnVideoPlaybackStopped(fn func(bot *Bot, request *model.VideoPlayerEventRequest))](#Bot.OnVideoPlaybackStopped)
  * [func (this *Bot) OnVideoPlayerScheduledStopReached(fn func(bot *Bot, request *model.VideoPlayerEventRequest))](#Bot.OnVideoPlayerScheduledStopReached)
  * [func (this *Bot) OnVideoRrogressReportIntevalElapsed(fn func(bot *Bot, request *model.VideoPlayerEventRequest))](#Bot.OnVideoRrogressReportIntevalElapsed)
  * [func (this *Bot) Run() string](#Bot.Run)


#### <a name="pkg-files">Package files</a>
[application.go](/src/github.com/dueros/bot-sdk-go/bot/application.go) [audio_player.go](/src/github.com/dueros/bot-sdk-go/bot/audio_player.go) [bot.go](/src/github.com/dueros/bot-sdk-go/bot/bot.go) [event_handler.go](/src/github.com/dueros/bot-sdk-go/bot/event_handler.go) [video_player.go](/src/github.com/dueros/bot-sdk-go/bot/video_player.go) 






## <a name="Bot">type</a> [Bot](/src/target/bot.go?s=98:1126#L10)
``` go
type Bot struct {
    Request  interface{}     // 对当前request的封装，需要在使用时断言，判断当前的类型
    Session  *model.Session  // 对session的封装
    Response *model.Response // 对技能返回的封装
    // contains filtered or unexported fields
}
```
技能基础类







### <a name="NewBot">func</a> [NewBot](/src/target/bot.go?s=1128:1163#L22)
``` go
func NewBot(rawRequest string) *Bot
```




### <a name="Bot.AddDefaultEventListener">func</a> (\*Bot) [AddDefaultEventListener](/src/target/bot.go?s=2186:2266#L52)
``` go
func (this *Bot) AddDefaultEventListener(fn func(bot *Bot, request interface{}))
```
添加事件默认处理函数
比如，在播放视频时，技能会收到各种事件的上报，如果不想一一处理可以使用这个来添加处理




### <a name="Bot.AddEventListener">func</a> (\*Bot) [AddEventListener](/src/target/bot.go?s=1864:1955#L44)
``` go
func (this *Bot) AddEventListener(eventName string, fn func(bot *Bot, request interface{}))
```
添加对事件的处理函数




### <a name="Bot.AddIntentHandler">func</a> (\*Bot) [AddIntentHandler](/src/target/bot.go?s=1659:1760#L37)
``` go
func (this *Bot) AddIntentHandler(intentName string, fn func(bot *Bot, request *model.IntentRequest))
```
添加对intent的处理函数




### <a name="Bot.OnAudioPlaybackFinished">func</a> (\*Bot) [OnAudioPlaybackFinished](/src/target/audio_player.go?s=896:995#L28)
``` go
func (this *Bot) OnAudioPlaybackFinished(fn func(bot *Bot, request *model.AudioPlayerEventRequest))
```
音频播放完成事件




### <a name="Bot.OnAudioPlaybackNearlyFinished">func</a> (\*Bot) [OnAudioPlaybackNearlyFinished](/src/target/audio_player.go?s=1212:1317#L36)
``` go
func (this *Bot) OnAudioPlaybackNearlyFinished(fn func(bot *Bot, request *model.AudioPlayerEventRequest))
```
音频快要播放结束上报的事件




### <a name="Bot.OnAudioPlaybackStarted">func</a> (\*Bot) [OnAudioPlaybackStarted](/src/target/audio_player.go?s=298:396#L12)
``` go
func (this *Bot) OnAudioPlaybackStarted(fn func(bot *Bot, request *model.AudioPlayerEventRequest))
```
音频开始播放事件




### <a name="Bot.OnAudioPlaybackStopped">func</a> (\*Bot) [OnAudioPlaybackStopped](/src/target/audio_player.go?s=597:695#L20)
``` go
func (this *Bot) OnAudioPlaybackStopped(fn func(bot *Bot, request *model.AudioPlayerEventRequest))
```
音频停止播放事件




### <a name="Bot.OnAudioRrogressReportIntevalElapsed">func</a> (\*Bot) [OnAudioRrogressReportIntevalElapsed](/src/target/audio_player.go?s=1532:1643#L44)
``` go
func (this *Bot) OnAudioRrogressReportIntevalElapsed(fn func(bot *Bot, request *model.AudioPlayerEventRequest))
```
音频周期上报播放进度




### <a name="Bot.OnDisplayElementSelected">func</a> (\*Bot) [OnDisplayElementSelected](/src/target/event_handler.go?s=270:359#L10)
``` go
func (this *Bot) OnDisplayElementSelected(fn func(bot *Bot, request *model.EventRequest))
```
ListTemplate 列表选择事件
Display.ButtonClicked 事件
<a href="https://dueros.baidu.com/didp/doc/dueros-bot-platform/dbp-custom/display-template_markdown#Display.ElementSelected%E4%BA%8B%E4%BB%B6">https://dueros.baidu.com/didp/doc/dueros-bot-platform/dbp-custom/display-template_markdown#Display.ElementSelected%E4%BA%8B%E4%BB%B6</a>




### <a name="Bot.OnLaunchRequest">func</a> (\*Bot) [OnLaunchRequest](/src/target/bot.go?s=2331:2412#L57)
``` go
func (this *Bot) OnLaunchRequest(fn func(bot *Bot, request *model.LaunchRequest))
```
打开技能时的处理




### <a name="Bot.OnLinkAccountSuccessed">func</a> (\*Bot) [OnLinkAccountSuccessed](/src/target/event_handler.go?s=956:1043#L38)
``` go
func (this *Bot) OnLinkAccountSuccessed(fn func(bot *Bot, request *model.EventRequest))
```
LinkAccountSucceeded 事件
```javascript
{


	"type": "Connections.Response",
	"name": "LinkAccountSucceeded",
	"requestId": "{{STRING}}",
	"timestamp": {{INT32}},
	"token": "{{STRING}}"

}
```




### <a name="Bot.OnScreenLinkClicked">func</a> (\*Bot) [OnScreenLinkClicked](/src/target/event_handler.go?s=1501:1585#L54)
``` go
func (this *Bot) OnScreenLinkClicked(fn func(bot *Bot, request *model.EventRequest))
```
Screen.LinkClicked事件
<a href="https://dueros.baidu.com/didp/doc/dueros-bot-platform/dbp-custom/cards_markdown#Screen.LinkClicked%E4%BA%8B%E4%BB%B6">https://dueros.baidu.com/didp/doc/dueros-bot-platform/dbp-custom/cards_markdown#Screen.LinkClicked%E4%BA%8B%E4%BB%B6</a>
{


	"type": "Screen.LinkClicked",
	"url": "{{STRING}}",
	"requestId": "{{STRING}}",
	"timestamp": {{INT32}}
	"token": "{{STRING}}"

}




### <a name="Bot.OnSessionEndedRequest">func</a> (\*Bot) [OnSessionEndedRequest](/src/target/bot.go?s=2598:2691#L63)
``` go
func (this *Bot) OnSessionEndedRequest(fn func(bot *Bot, request *model.SessionEndedRequest))
```
技能关闭的处理，比如可以做一些清理的工作
TIP: 根据协议，技能关闭返回的结果，DuerOS不会返回给用户。




### <a name="Bot.OnVideoPlaybackFinished">func</a> (\*Bot) [OnVideoPlaybackFinished](/src/target/video_player.go?s=896:995#L28)
``` go
func (this *Bot) OnVideoPlaybackFinished(fn func(bot *Bot, request *model.VideoPlayerEventRequest))
```
视频播放完成事件




### <a name="Bot.OnVideoPlaybackNearlyFinished">func</a> (\*Bot) [OnVideoPlaybackNearlyFinished](/src/target/video_player.go?s=1212:1317#L36)
``` go
func (this *Bot) OnVideoPlaybackNearlyFinished(fn func(bot *Bot, request *model.VideoPlayerEventRequest))
```
视频快要播放结束上报的事件




### <a name="Bot.OnVideoPlaybackStarted">func</a> (\*Bot) [OnVideoPlaybackStarted](/src/target/video_player.go?s=298:396#L12)
``` go
func (this *Bot) OnVideoPlaybackStarted(fn func(bot *Bot, request *model.VideoPlayerEventRequest))
```
视频开始播放事件




### <a name="Bot.OnVideoPlaybackStopped">func</a> (\*Bot) [OnVideoPlaybackStopped](/src/target/video_player.go?s=597:695#L20)
``` go
func (this *Bot) OnVideoPlaybackStopped(fn func(bot *Bot, request *model.VideoPlayerEventRequest))
```
视频停止播放事件




### <a name="Bot.OnVideoPlayerScheduledStopReached">func</a> (\*Bot) [OnVideoPlayerScheduledStopReached](/src/target/video_player.go?s=1863:1972#L52)
``` go
func (this *Bot) OnVideoPlayerScheduledStopReached(fn func(bot *Bot, request *model.VideoPlayerEventRequest))
```
视频自动暂停后上报




### <a name="Bot.OnVideoRrogressReportIntevalElapsed">func</a> (\*Bot) [OnVideoRrogressReportIntevalElapsed](/src/target/video_player.go?s=1532:1643#L44)
``` go
func (this *Bot) OnVideoRrogressReportIntevalElapsed(fn func(bot *Bot, request *model.VideoPlayerEventRequest))
```
视频周期上报播放进度




### <a name="Bot.Run">func</a> (\*Bot) [Run](/src/target/bot.go?s=2735:2764#L67)
``` go
func (this *Bot) Run() string
```







- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
