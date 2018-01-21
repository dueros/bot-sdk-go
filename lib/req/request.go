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
	if requestType == INENT_REQUEST {
		request := IntentRequest{}
		request.Type = requestType
		if err := json.Unmarshal(jsonBlob, &request.Data); err != nil {
			fmt.Println(err)
		}
		request.Nlu = nlu.NewDialog(request.Data.Request)
		return request
	} else if requestType == LAUNCH_REQUEST {
		request := LaunchRequest{}
		request.Type = requestType
		if err := json.Unmarshal(jsonBlob, &request.Data); err != nil {
			fmt.Println(err)
		}
		return request
	} else if requestType == SESSION_ENDED_REQUEST {
		request := SessionEndedRequest{}
		request.Type = requestType
		if err := json.Unmarshal(jsonBlob, &request.Data); err != nil {
			fmt.Println(err)
		}
		return request
	} else {
		request := EventRequest{}
		request.Type = requestType
		if err := json.Unmarshal(jsonBlob, &request.Data); err != nil {
			fmt.Println(err)
		}
		return request

	}
	return false
}

func (this *Request) GetDeviceId() string {
	return ""
}

//func (this *Request) getAudioPlayerContext() AudioPlayerContext{}

func (this *Request) GetAccessToken() string {
	return ""
}

func (this *Request) GetTimestamp() string {
	return ""
}

func (this *Request) GetBotId() string {
	return ""
}
