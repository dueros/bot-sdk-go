package main

import (
	dueros "github.com/dueros/bot-sdk-go/bot"
	"github.com/dueros/bot-sdk-go/bot/card"
	"github.com/dueros/bot-sdk-go/bot/directive/audio_player"
	"github.com/dueros/bot-sdk-go/bot/model"
)

func main() {

	bot := dueros.NewBot()                        // 创建bot 只创建一次
	bot.AddDefaultEventListener(defaultRequest)   // 添加默认事件的监听
	bot.AddIntentHandler("start_play", startPlay) // 添加开始播放的意图
	bot.OnLaunchRequest(launchRequest)            // 添加启动意图

	application := dueros.Application{AppId: "9f8f898b-2e91-a7ca-b25d-4aa9071f0470", Handler: bot.Handler} // 添加web容器
	application.Start(":8080")                                                                             // 启动服务器
}

func launchRequest(bot *dueros.Bot, request *model.LaunchRequest) {
	txtCard := card.NewTextCard("欢迎使用，请说开始测试")
	bot.Response.Tell("欢迎使用，请说开始测试")
	bot.Response.DisplayCard(txtCard)
}

func startPlay(bot *dueros.Bot, request *model.IntentRequest) {
	mp3 := audio_player.NewPlayDirective("http://baobaozhidao.bj.bcebos.com/beilehu%2F%E9%9D%92%E8%9B%99%E4%B9%90%E9%98%9F-%E5%B0%8F%E8%B7%B3%E8%9B%99.mp3")
	info := audio_player.NewPlayerInfo()
	info.SetTitle("小跳蛙")
	info.SetTitleSubtext1("xiaotiaowa")
	mp3.SetPlayerInfo(info)
	bot.Response.Command(mp3)
}

func defaultRequest(bot *dueros.Bot, request interface{}) {
	bot.Response.HoldOn()
	bot.Response.CloseMicrophone()
}
