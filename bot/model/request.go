package model

import (
	"encoding/json"
	"github.com/dueros/bot-sdk-go/bot/data"
	"log"
	"time"
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
	Dialog *Dialog
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

func (this *IntentRequest) GetIntentName() (string, bool) {
	return this.Dialog.GetIntentName()
}

func (this *IntentRequest) IsDialogStateCompleted() bool {
	return true
}

func (this *IntentRequest) GetQuery() string {
	query, _ := this.Dialog.GetQuery()
	return query
}

func (this *Request) GetUserId() string {
	return this.Common.Context.System.User.UserId
}

func (this *Request) GetDeviceId() string {
	return this.Common.Context.System.Device.DeviceId
}

func (this *Request) getAudioPlayerContext() data.AudioPlayerContext {
	return this.Common.Context.AudioPlayer
}

func (this *Request) GetAccessToken() string {
	return this.Common.Context.System.User.AccessToken
}

func (this *Request) GetTimestamp() string {
	return this.Common.Request.Timestamp
}

func (this *Request) GetRequestId() string {
	return this.Common.Request.RequestId
}

func (this *Request) GetBotId() string {
	return this.Common.Context.System.Application.ApplicationId
}

func (this *Request) VerifyTimestamp() bool {
	reqTimestamp, _ := time.Parse("2006-01-02T15:04:05Z", this.GetTimestamp())

	if time.Since(reqTimestamp) < time.Duration(180)*time.Second {
		return true
	}
	return false
}

func (this *Request) VerifyBotID(myBotID string) bool {
	if this.GetBotId() == myBotID {
		return true
	}
	return false
}

func getType(rawData string) string {
	jsonBlob := []byte(rawData)
	d := data.LaunchRequest{}

	if err := json.Unmarshal(jsonBlob, &d); err != nil {
		log.Println(err)
	}

	return d.Request.Type
}

func GetSessionData(rawData string) data.Session {
	jsonBlob := []byte(rawData)
	common := data.RequestPart{}
	if err := json.Unmarshal(jsonBlob, &common); err != nil {
		log.Println(err)
	}

	return common.Session
}

func NewRequest(rawData string) interface{} {
	requestType := getType(rawData)

	jsonBlob := []byte(rawData)

	common := data.RequestPart{}
	if err := json.Unmarshal(jsonBlob, &common); err != nil {
		log.Println(err)
		return false
	}

	if requestType == INENT_REQUEST {
		request := IntentRequest{}
		request.Type = requestType
		request.Common = common
		if err := json.Unmarshal(jsonBlob, &request.Data); err != nil {
			log.Println(err)
			return false
		}
		request.Dialog = NewDialog(request.Data.Request)

		return request
	} else if requestType == LAUNCH_REQUEST {
		request := LaunchRequest{}
		request.Type = requestType
		request.Common = common
		if err := json.Unmarshal(jsonBlob, &request.Data); err != nil {
			log.Println(err)
			return false
		}
		return request
	} else if requestType == SESSION_ENDED_REQUEST {
		request := SessionEndedRequest{}
		request.Type = requestType
		request.Common = common
		if err := json.Unmarshal(jsonBlob, &request.Data); err != nil {
			log.Println(err)
			return false
		}
		return request
	} else {
		request := EventRequest{}
		request.Type = requestType
		request.Common = common
		if err := json.Unmarshal(jsonBlob, &request.Data); err != nil {
			log.Println(err)
			return false
		}
		return request

	}
	return false
}
