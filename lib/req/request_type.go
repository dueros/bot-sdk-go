package req

import (
	"../data"
	"../nlu"
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

type Requset struct {
	Type string
}

type IntentRequest struct {
	Data data.IntentRequest
	Nlu  *nlu.Dialog
	Requset
}

type LaunchRequest struct {
	Data data.LaunchRequest
	Requset
}

type SessionEndedRequest struct {
	Data data.SessionEndedRequest
	Requset
}

type EventRequest struct {
	Data data.EventRequest
	Requset
}
