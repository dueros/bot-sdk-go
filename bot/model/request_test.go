package model

import (
	"fmt"
	//"log"
	"reflect"
	"strings"
	"testing"

	"github.com/dueros/bot-sdk-go/bot/util"
)

func TestGetUserId(t *testing.T) {
	functionTest("GetUserId", "4541", t)
}

func TestGetDeviceId(t *testing.T) {
	functionTest("GetDeviceId", "deviceId", t)
}

func TestGetTimestamp(t *testing.T) {
	functionTest("GetTimestamp", 1531750145, t)
}

func functionTest(funcName string, value interface{}, t *testing.T) {
	files := map[string]string{
		"IntentRequest":           "../test/intent-request.json",
		"LaunchRequest":           "../test/launch.json",
		"AudioPlayerEventRequest": "../test/audio-player-event.json",
		"VideoPlayerEventRequest": "../test/video-player-event.json",
	}

	for key, path := range files {
		var ret interface{}
		req := getRequest(path)
		split := strings.Split(reflect.TypeOf(req).String(), ".")
		reqType := split[len(split)-1]
		if reqType != key {
			t.Error(fmt.Sprintf("%s NewRequest not %s", path, funcName))
		}

		switch request := req.(type) {
		case LaunchRequest:
			ret = reflect.ValueOf(&request).MethodByName(funcName).Call(make([]reflect.Value, 0))[0]
		case IntentRequest:
			ret = reflect.ValueOf(&request).MethodByName(funcName).Call(make([]reflect.Value, 0))[0]
		case AudioPlayerEventRequest:
			ret = reflect.ValueOf(&request).MethodByName(funcName).Call(make([]reflect.Value, 0))[0]
		case VideoPlayerEventRequest:
			ret = reflect.ValueOf(&request).MethodByName(funcName).Call(make([]reflect.Value, 0))[0]
		}

		if !reflect.DeepEqual(ret.(reflect.Value).Interface(), value) {
			t.Error(fmt.Sprintf("%s error:%s ", funcName, key))
		}
	}
}

func getRequest(path string) interface{} {
	body, _ := util.ReadFileAll(path)
	rawRequest := string(body)
	request := NewRequest(rawRequest)

	return request
}
