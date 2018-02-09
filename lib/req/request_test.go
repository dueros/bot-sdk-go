package req

import "testing"

var rawData = `
		{
		  "version": "v2.0",
		  "session": {
		    "new": true,
		    "sessionId": "b818c84f-fa79-4ddc-8677-50b7a37c7764"
		  },
		  "context": {
		    "System": {
		      "user": {
		        "userId": "4541",
		        "accessToken": "21.2ce3445b90247e98ebea7c83f62f1425.2592000.1518757600.135166122-9943577",
		        "userInfo": {
		          "account": []
		        }
		      },
		      "application": {
		        "applicationId": "9c99871e-745f-9247-96b4-51b3dce8557f"
		      },
		      "device": {
		        "deviceId": "f56846491c8a034fc6497db4d1a7f96d",
		        "supportedInterfaces": {
		          "VoiceInput": [],
		          "VoiceOutput": [],
		          "AudioPlayer": []
		        }
		      }
		    }
		  },
		  "request": {
		    "query": {
		      "type": "TEXT",
		      "original": "查个税"
		    },
		    "dialogState": "STARTED",
		    "intents": [
		      {
		        "name": "test.inquiry888",
		        "confirmationStatus": "NONE",
		        "slots": {}
		      }
		    ],
		    "type": "IntentRequest",
		    "requestId": "3ea82668c5b44d4ab8e3414805a2d648_0",
		    "timestamp": "1516287724"
		  }
		}
		`

func TestNewRequest(t *testing.T) {
	req := NewRequest(rawData)

	request, ok := req.(IntentRequest)

	if !ok {
		t.Errorf("InentRequest Type Error")
	}

	if request.Common.Version != "v2.0" {
		t.Errorf("New Request Error")
	}
}
