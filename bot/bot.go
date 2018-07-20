package bot

import (
	"reflect"

	"github.com/dueros/bot-sdk-go/bot/model"
)

// 技能基础类
type Bot struct {
	intentHandler map[string]func(bot *Bot, request *model.IntentRequest) // 针对intent requset不同intent的处理函数
	//eventHandler               map[string]func(bot *Bot, request *model.EventRequest)  // 针对事件的处理函数
	eventHandler               map[string]func(bot *Bot, request interface{}) // 针对事件的处理函数
	defaultEventHandler        func(bot *Bot, request interface{})
	launchRequestHandler       func(bot *Bot, request *model.LaunchRequest)       // 针对技能打开的处理函数
	sessionEndedRequestHandler func(bot *Bot, request *model.SessionEndedRequest) // 针对技能关闭的处理函数
	Request                    interface{}                                        // 对当前request的封装，需要在使用时断言，判断当前的类型
	Session                    *model.Session                                     // 对session的封装
	Response                   *model.Response                                    // 对技能返回的封装
}

func NewBot(rawRequest string) *Bot {
	request := model.NewRequest(rawRequest)
	session := model.NewSession(model.GetSessionData(rawRequest))

	return &Bot{
		intentHandler: make(map[string]func(bot *Bot, request *model.IntentRequest)),
		//eventHandler:  make(map[string]func(bot *Bot, request *model.EventRequest)),
		eventHandler: make(map[string]func(bot *Bot, request interface{})),
		Request:      request,
		Session:      session,
		Response:     model.NewResponse(session, request),
	}
}

// 添加对intent的处理函数
func (this *Bot) AddIntentHandler(intentName string, fn func(bot *Bot, request *model.IntentRequest)) {
	if intentName != "" {
		this.intentHandler[intentName] = fn
	}
}

// 添加对事件的处理函数
func (this *Bot) AddEventListener(eventName string, fn func(bot *Bot, request interface{})) {
	if eventName != "" {
		this.eventHandler[eventName] = fn
	}
}

// 添加事件默认处理函数
// 比如，在播放视频时，技能会收到各种事件的上报，如果不想一一处理可以使用这个来添加处理
func (this *Bot) AddDefaultEventListener(fn func(bot *Bot, request interface{})) {
	this.defaultEventHandler = fn
}

// 打开技能时的处理
func (this *Bot) OnLaunchRequest(fn func(bot *Bot, request *model.LaunchRequest)) {
	this.launchRequestHandler = fn
}

// 技能关闭的处理，比如可以做一些清理的工作
// TIP: 根据协议，技能关闭返回的结果，DuerOS不会返回给用户。
func (this *Bot) OnSessionEndedRequest(fn func(bot *Bot, request *model.SessionEndedRequest)) {
	this.sessionEndedRequestHandler = fn
}

func (this *Bot) Run() string {
	this.dispatch()

	return this.Response.Build()
}

func (this *Bot) dispatch() {
	switch request := this.Request.(type) {
	case model.IntentRequest:
		this.processIntentHandler(request)
		return
	case model.LaunchRequest:
		this.processLaunchHandler(request)
		return
	case model.SessionEndedRequest:
		this.processSessionEndedHandler(request)
		return
	}
	this.processEventHandler(this.Request)
}

func (this *Bot) processLaunchHandler(request model.LaunchRequest) {
	if this.launchRequestHandler != nil {
		this.launchRequestHandler(this, &request)
	}
}

func (this *Bot) processSessionEndedHandler(request model.SessionEndedRequest) {
	if this.sessionEndedRequestHandler != nil {
		this.sessionEndedRequestHandler(this, &request)
	}
}

func (this *Bot) processIntentHandler(request model.IntentRequest) {
	intentName, _ := request.GetIntentName()
	fn, ok := this.intentHandler[intentName]

	if ok {
		fn(this, &request)
		return
	}
}

func (this *Bot) processEventHandler(req interface{}) {
	rVal := reflect.ValueOf(req)
	eventType := rVal.FieldByName("Type").Interface().(string)

	fn, ok := this.eventHandler[eventType]

	if ok {
		fn(this, req)
		return
	}

	this.defaultEventHandler(this, req)
}
