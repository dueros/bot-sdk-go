package req

import (
	"../data"
	"../nlu"
	"encoding/json"
	"fmt"
)

func getType(rawData string) string {
	jsonBlob := []byte(rawData)
	d := data.LaunchRequest{}

	if err := json.Unmarshal(jsonBlob, &d); err != nil {
		fmt.Println(err)
	}

	return d.Request.Type
}

func NewRequest(rawData string) interface{} {
	requestType := getType(rawData)

	jsonBlob := []byte(rawData)

	common := data.RequestPart{}
	if err := json.Unmarshal(jsonBlob, &common); err != nil {
		fmt.Println(err)
	}

	if requestType == INENT_REQUEST {
		request := IntentRequest{}
		request.Type = requestType
		request.Common = common
		if err := json.Unmarshal(jsonBlob, &request.Data); err != nil {
			fmt.Println(err)
		}
		request.Dialog = nlu.NewDialog(request.Data.Request)
		return request
	} else if requestType == LAUNCH_REQUEST {
		request := LaunchRequest{}
		request.Type = requestType
		request.Common = common
		if err := json.Unmarshal(jsonBlob, &request.Data); err != nil {
			fmt.Println(err)
		}
		return request
	} else if requestType == SESSION_ENDED_REQUEST {
		request := SessionEndedRequest{}
		request.Type = requestType
		request.Common = common
		if err := json.Unmarshal(jsonBlob, &request.Data); err != nil {
			fmt.Println(err)
		}
		return request
	} else {
		request := EventRequest{}
		request.Type = requestType
		request.Common = common
		if err := json.Unmarshal(jsonBlob, &request.Data); err != nil {
			fmt.Println(err)
		}
		return request

	}
	return false
}

func (this *Request) GetSession() data.Session {
	return this.Common.Session
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
