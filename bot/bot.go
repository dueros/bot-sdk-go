package bot

import (
	"github.com/dueros/bot-sdk-go/bot/model"
)

// 技能
type Bot struct {
	intentHandler              map[string]func(bot *Bot, request *model.IntentRequest) // 针对intent requset不同intent的处理函数
	eventHandler               map[string]func(bot *Bot, request *model.EventRequest)  // 针对事件的处理函数
	launchRequestHandler       func(bot *Bot, request *model.LaunchRequest)            // 针对技能打开的处理函数
	sessionEndedRequestHandler func(bot *Bot, request *model.SessionEndedRequest)      // 针对技能关闭的处理函数
	Request                    interface{}                                             // 对当前request的封装，需要在使用时断言，判断当前的类型
	Session                    *model.Session                                          // 对session的封装
	Response                   *model.Response                                         // 对技能返回的封装
}

func NewBot(rawRequest string) *Bot {
	request := model.NewRequest(rawRequest)
	session := getSession(request)

	return &Bot{
		intentHandler: make(map[string]func(bot *Bot, request *model.IntentRequest)),
		eventHandler:  make(map[string]func(bot *Bot, request *model.EventRequest)),
		Request:       request,
		Session:       session,
		Response:      model.NewResponse(session, request),
	}
}

func getSession(request interface{}) *model.Session {
	r, _ := request.(model.Request)
	return model.NewSession(r.GetSession())
}

// 添加对intent的处理函数
func (this *Bot) AddIntentHandler(intentName string, fn func(bot *Bot, request *model.IntentRequest)) {
	if intentName != "" {
		this.intentHandler[intentName] = fn
	}
}

// 添加对事件的处理函数
func (this *Bot) AddEventListener(eventName string, fn func(bot *Bot, request *model.EventRequest)) {
	if eventName != "" {
		this.eventHandler[eventName] = fn
	}
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

func (this *Bot) AskSlot(speech string, slot string) *Bot {

	return this
}

func (this *Bot) Run() string {
	this.dispatch()

	return this.Response.Build()
}

func (this *Bot) dispatch() {
	switch request := this.Request.(type) {
	case model.IntentRequest:
		this.processIntentHandler(request)
	case model.LaunchRequest:
	case model.SessionEndedRequest:
	case model.EventRequest:
		this.processEventHandler(request)
	}
}

func (this *Bot) processIntentHandler(request model.IntentRequest) {
	intentName, _ := request.GetIntentName()
	fn, ok := this.intentHandler[intentName]

	if ok {
		fn(this, &request)
	}
}

func (this *Bot) processEventHandler(request model.EventRequest) {
	fn, ok := this.eventHandler[request.Type]

	if ok {
		fn(this, &request)
	}
}
