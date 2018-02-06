package req

import (
	"github.com/dueros/bot-sdk-go/lib/data"
	"github.com/dueros/bot-sdk-go/lib/nlu"
)

const (
	INENT_REQUEST         = "IntentRequest"
	LAUNCH_REQUEST        = "LaunchRequest"
	SESSION_ENDED_REQUEST = "SessionEndedRequest"

	AUDIO_PLAYER_PLAYBACK_STARTED                 = "AudioPlayer.PlaybackStarted"
	AUDIO_PLAYER_PLAYBACK_STOPPED                 = "AudioPlayer.PlaybackStopped"
	AUDIO_PLAYER_PLAYBACK_FINISHED                = "AudioPlayer.PlaybackFinished"
	AUDIO_PLAYER_PLAYBACK_NEARLY_FINISHED         = "AudioPlayer.PlaybackNearlyFinished"
	AUDIO_PLAYER_PROGRESS_REPORT_INTERVAL_ELAPSED = "AudioPlayer.ProgressReportIntervalElapsed"
)

type Request struct {
	Type   string
	Common data.RequestPart
}

type IntentRequest struct {
	Data   data.IntentRequest
	Dialog *nlu.Dialog
	Request
}

type LaunchRequest struct {
	Data data.LaunchRequest
	Request
}

type SessionEndedRequest struct {
	Data data.SessionEndedRequest
	Request
}

type EventRequest struct {
	Data data.EventRequest
	Request
}
