package lib

import (
	"./data"
	"encoding/json"
	"fmt"
	"regexp"
)

type Response struct {
	data map[string]interface{}
}

func NewResponse(session *Session, request interface{}) *Response {
	data := make(map[string]interface{})
	return &Response{
		data: data,
	}
}

func (this *Response) Ask(speech string) *Response {
	this.Tell(speech)
	this.HoldOn()
	return this
}

func (this *Response) Tell(speech string) *Response {
	this.data["outputSpeech"] = this.formatSpeech(speech)
	return this
}

func (this *Response) DisplayCard(card interface{}) *Response {
	this.data["card"] = card

	return this
}

func (this *Response) Command(directive interface{}) *Response {
	_, ok := this.data["directives"]
	if !ok {
		this.data["directives"] = make([]interface{}, 0)
	}

	directives, ok := this.data["directives"].([]interface{})
	directives = append(directives, directive)

	this.data["directives"] = directives

	return this
}

func (this *Response) HoldOn() *Response {
	this.data["shouldEndSession"] = false
	return this
}

func (this *Response) Build() string {
	ret := map[string]interface{}{
		"version":  "2.0",
		"response": this.data,
	}

	res2B, _ := json.Marshal(ret)

	fmt.Println(string(res2B))
	return string(res2B)
}

func (this *Response) formatSpeech(speech string) data.Speech {
	match, _ := regexp.MatchString("^<speak>", speech)

	if match {
		return data.Speech{
			Type: "SSML",
			Ssml: speech,
		}
	}

	return data.Speech{
		Type: "PlainText",
		Text: speech,
	}
}
