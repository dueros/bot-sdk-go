package audio_player

import (
	"github.com/dueros/bot-sdk-go/bot/directive"
	"log"
)

var behavior = map[string]string{
	"REPLACE_ALL":      "REPLACE_ALL",
	"REPLACE_ENQUEUED": "REPLACE_ENQUEUED",
	"ENQUEUE":          "ENQUEUE",
}

var audioFormat = map[string]string{
	"AUDIO_MP3":  "AUDIO_MP3",
	"AUDIO_M3U8": "AUDIO_M3U8",
	"AUDIO_M4A":  "AUDIO_M4A",
}

func checkBehavior(bhv string) string {
	data, ok := behavior[bhv]
	if !ok {
		log.Fatal("behavior error")
	}
	return data
}

func checkformat(fmt string) string {
	data, ok := audioFormat[fmt]
	if !ok {
		log.Fatal("audioFormat error")
	}
	return data
}

type PlayDirective struct {
	directive.BaseDirective
	PlayBehavior string `json:"playBehavior"`
	AudioItem    struct {
		Stream struct {
			Url                      string `json:"url"`
			StreamFormat             string `json:"streamFormat"`
			OffsetInMilliSeconds     int    `json:"offsetInMilliSeconds"`
			Token                    string `json:"token"`
			ProgressReportIntervalMs int    `json:"progressReportIntervalMs"`
		} `json:"stream"`
	} `json:"audioItem"`
}

func NewPlayDirective(behavior, url, format string, offset, interval int) *PlayDirective {
	play := &PlayDirective{}
	play.Type = "AudioPlayer.Play"
	play.PlayBehavior = checkBehavior(behavior)
	play.AudioItem.Stream.Url = url
	play.AudioItem.Stream.StreamFormat = checkformat(format)
	play.AudioItem.Stream.OffsetInMilliSeconds = offset
	play.AudioItem.Stream.Token = play.GenToken()
	play.AudioItem.Stream.ProgressReportIntervalMs = interval
	return play
}
