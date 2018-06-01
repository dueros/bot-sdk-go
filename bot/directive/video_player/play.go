package video_player

import (
	"github.com/dueros/bot-sdk-go/bot/directive"
)

var behaviorMap = map[string]bool{
	ENQUEUE:          true,
	REPLACE_ALL:      true,
	REPLACE_ENQUEUED: true,
}

type PlayDirective struct {
	directive.BaseDirective
	PlayBehavior string `json:"playBehavior"`
	VideoItem    struct {
		Stream struct {
			Url                  string `json:"url"`
			OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
			ExpiryTime           string `json:"expiryTime,omitempty"`
			ProgressReport       struct {
				ProgressReportDelayInMilliseconds    int `json:"progressReportDelayInMilliseconds,omitempty"`
				ProgressReportIntervalInMilliseconds int `json:"progressReportIntervalInMilliseconds,omitempty"`
			} `json:"progressReport,omitempty"`
			Token                 string `json:"token"`
			ExpectedPreviousToken string `json:"expectedPreviousToken,omitempty"`
		} `json:"stream"`
	} `json:"VideoItem"`
}

func NewPlayDirective(url string) *PlayDirective {
	play := &PlayDirective{}
	play.Type = "VideoPlayer.Play"
	play.PlayBehavior = REPLACE_ALL
	play.VideoItem.Stream.Url = url
	play.VideoItem.Stream.OffsetInMilliseconds = 0
	play.VideoItem.Stream.Token = play.GenToken()
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
	this.VideoItem.Stream.Token = token
	return this
}

func (this *PlayDirective) GetToken(token string) string {
	return this.VideoItem.Stream.Token
}

func (this *PlayDirective) SetUrl(url string) *PlayDirective {
	this.VideoItem.Stream.Url = url
	return this
}

func (this *PlayDirective) SetOffsetInMilliseconds(milliseconds int) *PlayDirective {
	this.VideoItem.Stream.OffsetInMilliseconds = milliseconds
	return this
}

func (this *PlayDirective) SetExpiryTime(expiryTime string) *PlayDirective {
	this.VideoItem.Stream.ExpiryTime = expiryTime
	return this
}

func (this *PlayDirective) SetReportDelayInMs(reportDelayInMs int) *PlayDirective {
	this.VideoItem.Stream.ProgressReport.ProgressReportDelayInMilliseconds = reportDelayInMs
	return this
}

func (this *PlayDirective) SetReportIntervalInMs(reportIntervalInMs int) *PlayDirective {
	this.VideoItem.Stream.ProgressReport.ProgressReportIntervalInMilliseconds = reportIntervalInMs
	return this
}

func (this *PlayDirective) SetExpectedPreviousToken(expectedPreviousToken string) *PlayDirective {
	this.VideoItem.Stream.ExpectedPreviousToken = expectedPreviousToken
	return this
}
