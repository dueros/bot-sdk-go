# 度秘BOT SDK for GO
这是一个帮助开发Bot的SDK，针对php、java、node.js也有对应的版本.SDK封装了对session、nlu、result的处理，使用sdk可以避免由于BOT的协议的升级带来一些麻烦。这个SDK会与DuerOS的协议一起升级，会最大限度减少对您开发bot的影响。

## 通过bot-sdk可以快速的开发bot
我们的目标是通过使用bot-sdk，可以迅速的开发一个bot，而不必过多去关注DuerOS对Bot的复杂协议。我们提供如下功能：

* 封装了DuerOS的request和response
* 提供了session简化接口
* 提供了nlu简化接口
    * slot 操作
    * nlu理解交互接口（ask）
* 提供了多轮对话开发接口
* 提供了事件监听接口
* 展现模板、视频、音频指令的封装

## 安装、使用BOT SDK进行开发 
度秘BOT SDK使用执行如下命令进行安装：
```shell
go get github.com/dueros/bot-sdk-go
```

为了开始使用BOT SDK，你需要先新建一个main文件，比如文件名是main.go。

```javascript
import (
	dueros "github.com/dueros/bot-sdk-go/bot"
	"github.com/dueros/bot-sdk-go/bot/directive/display"
	"github.com/dueros/bot-sdk-go/bot/directive/display/template"
	"github.com/dueros/bot-sdk-go/bot/model"
)

func main() {

	application := dueros.Application{AppId: "you bot id", Handler: func(body string) string {
		bot := dueros.NewBot(body)

        // 意图处理函数
		bot.AddIntentHandler("test.inquiry888",
			func(this *dueros.Bot, request *model.IntentRequest) {
				template := template.NewBodyTemplate1()
				template.SetTitle("title")
				template.SetPlainContent("欢迎查询")
				template.SetBackgroundImageUrl("http://www.baidu.com")
				directive := display.NewRenderTemplate(template)

				this.Response.Ask("欢迎查询").Command(directive)
			})

        // 技能被打开的请求处理
		bot.OnLaunchRequest(func(this *dueros.Bot, request *model.LaunchRequest) {
			this.Response.Tell("欢迎使用查个税")
		})

		return bot.Run()
	}}
	application.Start(":8888")
}
```
下一步，我们处理意图。Bot-sdk提供了一个函数来handle这些意图。比如，查询个税，创建一个handler，在构造函数中添加：


这里`AddIntentHandler`可以用来建立(intent) => handler的映射，第一个参数是条件，如果满足则执行对应的回调函数(第二个参数)。
其中，this指向当前的Bot，通过Response来进行回复。`Ask`可以用来向用户澄清，`Command`可以用来返回一个指令。

## API 文档

* [Bot](doc/bot.md)
* [Application](doc/application.md)
* model
    * [Request](doc/model/request.md)
        * [IntentRequest](doc/model/request.md)
        * [LaunchRequest](doc/model/request.md)
        * [SessionEndedRequest](doc/model/request.md)
        * [EventRequest](doc/model/request.md)
    * [Response](doc/model/response.md)
    * [Dialog](doc/model/dialog.md)
    * [Intent](doc/model/intent.md)
    * [Session](doc/model/session.md)
    * [SSMLTextBuilder](doc/model/ssml-text-builder.md)
* [展现卡片(card)](doc/card/card.md)
* [展现模版(template)](doc/directive/display/render-template.md)
    * [template](doc/directive/display/template.md)
* [播放音频(audio)](doc/directive/audio-player.md)
* [播放视频(video)](doc/directive/video-player.md)
