package audio_player

import (
	"github.com/dueros/bot-sdk-go/bot/directive"
)

var behaviorMap = map[string]bool{
	ENQUEUE:          true,
	REPLACE_ALL:      true,
	REPLACE_ENQUEUED: true,
}

var audioFormatMap = map[string]bool{
	AUDIO_MP3:  true,
	AUDIO_M3U8: true,
	AUDIO_M4A:  true,
}

type PlayDirective struct {
	directive.BaseDirective
	PlayBehavior string `json:"playBehavior"`
	AudioItem    struct {
		Stream struct {
			Url                      string `json:"url"`
			StreamFormat             string `json:"streamFormat"`
			OffsetInMilliseconds     int    `json:"offsetInMilliSeconds"`
			Token                    string `json:"token"`
			ProgressReportIntervalMs int    `json:"progressReportIntervalMs,omitempty"`
		} `json:"stream"`
		PlayerInfo *PlayerInfo `json:"playerInfo,omitempty"`
	} `json:"audioItem"`
}

func NewPlayDirective(url string) *PlayDirective {
	play := &PlayDirective{}
	play.Type = "AudioPlayer.Play"
	play.PlayBehavior = REPLACE_ALL
	play.AudioItem.Stream.Url = url
	play.AudioItem.Stream.StreamFormat = AUDIO_MP3
	play.AudioItem.Stream.OffsetInMilliseconds = 0
	play.AudioItem.Stream.Token = play.GenToken()
	play.AudioItem.PlayerInfo = nil
	//play.AudioItem.Stream.ProgressReportIntervalMs = 20
	return play
}

func (this *PlayDirective) SetBehavior(behavior string) *PlayDirective {
	_, ok := behaviorMap[behavior]
	if ok {
		this.PlayBehavior = behavior
	}

	return this
}

func (this *PlayDirective) SetToken(token string) *PlayDirective {
	this.AudioItem.Stream.Token = token
	return this
}

func (this *PlayDirective) GetToken(token string) string {
	return this.AudioItem.Stream.Token
}

func (this *PlayDirective) SetUrl(url string) *PlayDirective {
	this.AudioItem.Stream.Url = url
	return this
}

func (this *PlayDirective) SetOffsetInMilliseconds(milliseconds int) *PlayDirective {
	this.AudioItem.Stream.OffsetInMilliseconds = milliseconds
	return this
}

func (this *PlayDirective) SetProgressReportIntervalMs(intervalMs int) *PlayDirective {
	this.AudioItem.Stream.ProgressReportIntervalMs = intervalMs
	return this
}

func (this *PlayDirective) SetStreamFormat(format string) *PlayDirective {
	_, ok := audioFormatMap[format]
	if ok {
		this.AudioItem.Stream.StreamFormat = format
	}
	return this
}

func (this *PlayDirective) SetPlayerInfo(playerInfo *PlayerInfo) *PlayDirective {
	this.AudioItem.PlayerInfo = playerInfo
	return this
}
