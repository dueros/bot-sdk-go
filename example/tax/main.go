package main

import (
	dueros "github.com/dueros/bot-sdk-go/bot"
	//"github.com/dueros/bot-sdk-go/bot/card"
	//"github.com/dueros/bot-sdk-go/bot/directive/audio_player"
	"github.com/dueros/bot-sdk-go/bot/directive/display"
	"github.com/dueros/bot-sdk-go/bot/directive/display/template"
	//"github.com/dueros/bot-sdk-go/bot/directive/video_player"
	"github.com/dueros/bot-sdk-go/bot/model"
)

func main() {

	application := dueros.Application{AppId: "17a64831-b273-8af5-f995-9a805e45cffe", Handler: func(body string) string {
		bot := dueros.NewBot(body)
		bot.AddIntentHandler("test.inquiry888",
			func(this *dueros.Bot, request *model.IntentRequest) {
				this.Session.SetAttribute("key1", "test")
				template := template.NewBodyTemplate1()
				template.SetTitle("hah")
				template.SetPlainContent("欢迎查询")
				template.SetBackgroundImageUrl("http://www.baidu.com")
				directive := display.NewRenderTemplate(template)
				//c := card.NewTextCard("欢迎查询")

				//audio := audio_player.NewPlayDirective("http://www.wg")
				//playInfo := audio_player.NewPlayerInfo()
				//playInfo.SetTitle(request.GetUserId())
				//audio.SetPlayerInfo(playInfo)

				//video := video_player.NewPlayDirective("http://wwgwgw")
				//video.SetReportIntervalInMs(20000)

				//request.Dialog.ConfirmIntent()
				this.Response.AskSlot("工资多少呢？", "salary").Command(directive)

				//ssmlBuilder := model.NewSSMLTextBuilder()
				//ssmlBuilder.AppendPlainSpeech("欢迎").AppendAudio("http://www.gasg.mp3").ApplyBackground("http://www.ga", true)

				//this.Response.Ask("欢迎查询").Command(directive)
				//this.Response.Ask("欢迎查询").Command(directive).Command(audio).Command(video)
				//this.Response.Ask(ssmlBuilder.Build()).Command(directive).Command(audio).Command(video)
			})

		bot.OnLaunchRequest(func(this *dueros.Bot, request *model.LaunchRequest) {
			this.Response.Tell("欢迎使用查个税")
		})

		return bot.Run()
	}}
	application.Start(":8888")
}
