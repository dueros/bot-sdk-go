package lib

//import "fmt"

type ResponseData struct {
	Card string
}

type Bot struct {
	intentHandler              map[string]func() ResponseData
	eventHandler               map[string]func() ResponseData
	launchRequestHandler       func() ResponseData
	sessionEndedRequestHandler func() ResponseData
	Request                    interface{}
}

func NewBot(rawRequest string) *Bot {
	return &Bot{
		intentHandler: make(map[string]func() ResponseData),
		eventHandler:  make(map[string]func() ResponseData),
		Request:       NewRequest(rawRequest),
	}
}

func (this *Bot) AddIntentHandler(intentName string, fn func() ResponseData) {
	if intentName != "" {
		this.intentHandler[intentName] = fn
	}
}

func (this *Bot) AddEventListener(eventName string, fn func() ResponseData) {
	if eventName != "" {
		this.eventHandler[eventName] = fn
	}
}

func (this *Bot) onLaunchRequest(fn func() ResponseData) {
	this.launchRequestHandler = fn
}

func (this *Bot) onSessionEndedRequest(fn func() ResponseData) {
	this.sessionEndedRequestHandler = fn
}

func (this *Bot) dispatch() string {
	return ""
}
