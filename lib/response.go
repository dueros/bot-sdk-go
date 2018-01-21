package lib

import (
	"./card"
	"encoding/json"
	"fmt"
)

type Response struct {
	data map[string]interface{}
}

func NewResponse() *Response {
	return &Response{
		data: make(map[string]interface{}),
	}
}

func (this *Response) Ask(speech string) *Response {

	this.HoldOn()
	return this
}

func (this *Response) Tell(speech string) *Response {

	return this
}

func (this *Response) DisplayCard() *Response {
	this.data["card"] = card.BaseCard{Type: "tct"}

	return this
}

func (this *Response) Command() *Response {

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
