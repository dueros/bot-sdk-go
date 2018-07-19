package model

import (
	"reflect"
	"testing"

	"github.com/dueros/bot-sdk-go/bot/data"
	"github.com/dueros/bot-sdk-go/bot/util"
	//"log"
)

func TestAsk(t *testing.T) {
	response := getResponse("../test/launch.json")
	response.Ask("ask")

	ret := response.GetData()
	shouldEndSession := ret["shouldEndSession"].(bool)

	if shouldEndSession != false {
		t.Error("Ask: shouldEndSession is not false")
	}

	speech := ret["outputSpeech"].(data.Speech)

	if !reflect.DeepEqual(speech, data.Speech{Type: "PlainText", Text: "ask"}) {
		t.Error("Ask: output speech is not ask")
	}

	t.Log("Ask: passed")
}

func TestAskSlot(t *testing.T) {
	response := getResponse("../test/intent-request.json")
	response.AskSlot("ask", "slot1")
	response.Build()

	ret := response.GetData()
	shouldEndSession := ret["shouldEndSession"].(bool)

	if shouldEndSession != false {
		t.Error("AskSlot: shouldEndSession is not false")
	}

	speech := ret["outputSpeech"].(data.Speech)

	if !reflect.DeepEqual(speech, data.Speech{Type: "PlainText", Text: "ask"}) {
		t.Error("AskSlot: output speech is not ask")
	}

	arr, ok := ret["directives"].([]interface{})
	directive, ok := arr[0].(*data.DialogDirective)

	if !ok {
		t.Error("AskSlot: no dialog.ElicitSlot")
	}

	if directive.Type != "Dialog.ElicitSlot" {
		t.Error("AskSlot: directive is not dialog.ElicitSlot ")
	}

	if directive.SlotToElicit != "slot1" {
		t.Error("AskSlot: directive.SlotToElicit is not slot1")
	}
}

func TestTell(t *testing.T) {
	response := getResponse("../test/launch.json")
	response.Tell("tell")

	ret := response.GetData()

	speech := ret["outputSpeech"].(data.Speech)

	if !reflect.DeepEqual(speech, data.Speech{Type: "PlainText", Text: "tell"}) {
		t.Error("Tell: output speech is not tell")
	}

	t.Log("Tell: passed")
}

func TestReprompt(t *testing.T) {
	response := getResponse("../test/launch.json")
	response.Reprompt("reprompt")

	ret := response.GetData()

	reprompt, ok := ret["reprompt"].(map[string]interface{})
	if !ok {
		t.Error("Reprompt: reprompt is not map structure")
	}

	outputSpeech, ok := reprompt["outputSpeech"]
	if !ok {
		t.Error("Reprompt: reprompt has no outputSpeech")
	}

	speech := outputSpeech.(data.Speech)

	if !reflect.DeepEqual(speech, data.Speech{Type: "PlainText", Text: "reprompt"}) {
		t.Error("Reprompt: output speech is not reprompt")
	}

	t.Log("Reprompt: passed")
}

func TestDisplayCard(t *testing.T) {}
func TestCommand(t *testing.T)     {}

func TestHoldOn(t *testing.T) {
	response := getResponse("../test/launch.json")
	response.HoldOn()

	ret := response.GetData()
	shouldEndSession := ret["shouldEndSession"].(bool)

	if shouldEndSession != false {
		t.Error("HoldOn: shouldEndSession is not false")
	}
}

func TestCloseMicrophone(t *testing.T) {
	response := getResponse("../test/launch.json")
	response.CloseMicrophone()

	ret := response.GetData()
	expectSpeech := ret["expectSpeech"].(bool)

	if expectSpeech != true {
		t.Error("CloseMicrophone: expectSpeech is not true")
	}
}

func getResponse(path string) *Response {
	body, _ := util.ReadFileAll(path)
	rawRequest := string(body)
	request := NewRequest(rawRequest)
	session := NewSession(GetSessionData(rawRequest))

	response := NewResponse(session, request)

	return response
}
