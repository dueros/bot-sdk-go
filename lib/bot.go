package lib

import (
	"github.com/dueros/bot-sdk-go/lib/req"
)

type Bot struct {
	intentHandler              map[string]func(bot *Bot, request *req.IntentRequest)
	eventHandler               map[string]func(bot *Bot, request *req.EventRequest)
	launchRequestHandler       func(bot *Bot, request *req.LaunchRequest)
	sessionEndedRequestHandler func(bot *Bot, request *req.SessionEndedRequest)
	Request                    interface{}
	Session                    *Session
	Response                   *Response
}

func NewBot(rawRequest string) *Bot {
	request := req.NewRequest(rawRequest)
	session := getSession(request)

	return &Bot{
		intentHandler: make(map[string]func(bot *Bot, request *req.IntentRequest)),
		eventHandler:  make(map[string]func(bot *Bot, request *req.EventRequest)),
		Request:       request,
		Session:       session,
		Response:      NewResponse(session, request),
	}
}

func getSession(request interface{}) *Session {
	r, _ := request.(req.Request)
	return NewSession(r.GetSession())
}

func (this *Bot) AddIntentHandler(intentName string, fn func(bot *Bot, request *req.IntentRequest)) {
	if intentName != "" {
		this.intentHandler[intentName] = fn
	}
}

func (this *Bot) AddEventListener(eventName string, fn func(bot *Bot, request *req.EventRequest)) {
	if eventName != "" {
		this.eventHandler[eventName] = fn
	}
}

func (this *Bot) OnLaunchRequest(fn func(bot *Bot, request *req.LaunchRequest)) {
	this.launchRequestHandler = fn
}

func (this *Bot) OnSessionEndedRequest(fn func(bot *Bot, request *req.SessionEndedRequest)) {
	this.sessionEndedRequestHandler = fn
}

func (this *Bot) AskSlot(speech string, slot string) *Bot {

	return this
}

func (this *Bot) Run() string {
	this.dispatch()

	return this.Response.Build()
}

func (this *Bot) dispatch() {
	switch request := this.Request.(type) {
	case req.IntentRequest:
		this.processIntentHandler(request)
	case req.LaunchRequest:
	case req.SessionEndedRequest:
	case req.EventRequest:
		this.processEventHandler(request)
	}
}

func (this *Bot) processIntentHandler(request req.IntentRequest) {
	intentName, _ := request.GetIntentName()
	fn, ok := this.intentHandler[intentName]

	if ok {
		fn(this, &request)
	}
}

func (this *Bot) processEventHandler(request req.EventRequest) {
	fn, ok := this.eventHandler[request.Type]

	if ok {
		fn(this, &request)
	}
}
