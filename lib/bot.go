package lib

import (
	"github.com/dueros/bot-sdk-go/lib/req"
)

// 技能
type Bot struct {
	intentHandler              map[string]func(bot *Bot, request *req.IntentRequest) // 针对intent requset不同intent的处理函数
	eventHandler               map[string]func(bot *Bot, request *req.EventRequest)  // 针对事件的处理函数
	launchRequestHandler       func(bot *Bot, request *req.LaunchRequest)            // 针对技能打开的处理函数
	sessionEndedRequestHandler func(bot *Bot, request *req.SessionEndedRequest)      // 针对技能关闭的处理函数
	Request                    interface{}                                           // 对当前request的封装，需要在使用时断言，判断当前的类型
	Session                    *Session                                              // 对session的封装
	Response                   *Response                                             // 对技能返回的封装
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

// 添加对intent的处理函数
func (this *Bot) AddIntentHandler(intentName string, fn func(bot *Bot, request *req.IntentRequest)) {
	if intentName != "" {
		this.intentHandler[intentName] = fn
	}
}

// 添加对事件的处理函数
func (this *Bot) AddEventListener(eventName string, fn func(bot *Bot, request *req.EventRequest)) {
	if eventName != "" {
		this.eventHandler[eventName] = fn
	}
}

// 打开技能时的处理
func (this *Bot) OnLaunchRequest(fn func(bot *Bot, request *req.LaunchRequest)) {
	this.launchRequestHandler = fn
}

// 技能关闭的处理，比如可以做一些清理的工作
// TIP: 根据协议，技能关闭返回的结果，DuerOS不会返回给用户。
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
