package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	dueros "github.com/dueros/bot-sdk-go/bot"
	"github.com/dueros/bot-sdk-go/bot/directive/display"
	"github.com/dueros/bot-sdk-go/bot/directive/display/template"
	"github.com/dueros/bot-sdk-go/bot/model"
)

var CODE_INIT = false
var CODE = make(map[string]City)

type City struct {
	Code string
	Name string
}

type CityList struct {
	Cities  []City
	DataVer int32
}

type Result struct {
	FlightNo    string
	Airname     string
	DDate       string
	Price       int
	Departure   string
	Destination string
}

type CitySlot struct {
	City string
}

func main() {

	bot := dueros.NewBot()                             // 创建bot 只创建一次
	bot.OnLaunchRequest(launchRequest)                 // 添加启动意图
	bot.AddIntentHandler("flightSearch", flightSearch) // 添加查询意图处理

	application := dueros.Application{AppId: "14a6ee9e-3367-f4e3-6698-d8df54c8a104", Handler: bot.Handler} // 添加web容器
	application.Start(":8880")
}

func launchRequest(this *dueros.Bot, request *model.LaunchRequest) {
	this.Response.HoldOn().Tell("欢迎使用机票查询服务，告诉我出发时间、离开城市、到达城市就可以查询最低的机票价格，比如：明天北京到上海的机票")
}

func flightSearch(this *dueros.Bot, request *model.IntentRequest) {
	if !request.IsDialogStateCompleted() && !request.IsSupportDisplay() {
		request.Dialog.Delegate()
		return
	}

	date := request.Dialog.GetSlotValue("date")
	departure := request.Dialog.GetSlotValue("departure")
	destination := request.Dialog.GetSlotValue("destination")

	if request.IsSupportDisplay() {
		if date == "" {
			text := "请问什么时间出发呢？"
			display := getResult(text, "机票助手")
			this.Response.AskSlot(text, "date").Reprompt(text).Command(display)
			return
		} else if departure == "" {
			text := "请问你的出发城市是哪里？"
			display := getResult(text, "机票助手")
			this.Response.AskSlot(text, "departure").Reprompt(text).Command(display)
			return
		} else if destination == "" {
			text := "请问你的目的城市是哪里？"
			display := getResult(text, "机票助手")
			this.Response.AskSlot(text, "destination").Reprompt(text).Command(display)
			return
		}
	}

	departureCity := getCity(departure)
	destinationCity := getCity(destination)

	departureCode := getCityCode(departureCity)
	destinationCode := getCityCode(destinationCity)

	result := getFlight(date, departureCode, destinationCode)

	if result != nil {
		result.Departure = departureCity
		result.Destination = destinationCity

		text := "价格" + strconv.Itoa(result.Price)
		this.Response.Tell(text)
		if request.IsSupportDisplay() {
			display := getResult(text, "机票助手")
			this.Response.Command(display)
		}
		return
	}

	this.Response.Tell("查询失败")
	return
}

func getCity(jsonStr string) string {
	city := CitySlot{}
	if err := json.Unmarshal([]byte(jsonStr), &city); err != nil {
		log.Fatal("city slot parse error")
		return ""
	}

	return city.City
}

func getFlight(date, departure, destination string) *Result {
	url := "http://sec-m.ctrip.com/restapi/soa2/13516/lowPriceCalendar"
	req := fmt.Sprintf(`{"head":{},"stype":1,"dcty":"%s","acty":"%s","start":"","end":"%s","flag":1}`, departure, destination, date)
	log.Println(req)

	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(req)))
	request.Header.Set("Content-type", "application/json")
	response, _ := client.Do(request)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		log.Println(string(body))

		d := struct {
			Prices []Result
		}{}

		if err := json.Unmarshal(body, &d); err != nil {
			log.Fatal("flight result parse error")
			return nil
		}

		for _, item := range d.Prices {
			if item.DDate == date {
				return &item
			}
		}
	}
	return nil
}

func getCityCode(city string) string {
	if !CODE_INIT {
		f, err := os.Open("city.json")
		defer f.Close()

		if err != nil {
			return ""
			log.Fatal("city.json file ready error")
		}

		d := CityList{}
		jsonBlob, _ := ioutil.ReadAll(f)
		if err := json.Unmarshal(jsonBlob, &d); err != nil {
			log.Fatal("json parse error")
			return ""
		}

		for _, cityItem := range d.Cities {
			CODE[cityItem.Name] = cityItem
		}
		CODE_INIT = true
	}

	return CODE[city].Code
}

func getResult(text string, title string) *display.RenderTemplate {
	template := template.NewBodyTemplate1()
	template.SetTitle(title)
	template.SetPlainContent(text)
	template.SetBackgroundImageUrl("http://img.taopic.com/uploads/allimg/120801/214828-120P10Z43585.jpg")

	return display.NewRenderTemplate(template)
}
