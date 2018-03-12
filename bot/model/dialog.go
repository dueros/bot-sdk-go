package model

import (
	"github.com/dueros/bot-sdk-go/bot/data"
)

type Dialog struct {
	data        data.IntentRequestBody
	Intents     []*Intent
	DialogState string
	directive   *data.DialogDirective
}

func NewDialog(request data.IntentRequestBody) *Dialog {
	length := len(request.Intents)
	intents := make([]*Intent, length)
	for i := 0; i < length; i++ {
		intents[i] = NewIntent(request.Intents[i])
	}
	return &Dialog{
		data:        request,
		Intents:     intents,
		DialogState: request.DialogState,
	}
}

func (this *Dialog) getIntent(index int) *Intent {
	l := len(this.Intents)
	if l > 0 && index < l {
		intent := this.Intents[index]
		return intent
	}
	return nil
}

func getIndex(index []int) int {
	i := 0
	if len(index) > 0 {
		i = index[0]
	}
	return i
}

// 获取用户请求的原始query
func (this *Dialog) GetQuery() (string, bool) {
	if this.data.Query.Type == "TEXT" {
		return this.data.Query.Original, true
	}
	return "", false
}

// 获取当前意图的名字
func (this *Dialog) GetIntentName() (string, bool) {
	intent := this.getIntent(0)
	if intent != nil {
		return intent.Name, true
	}
	return "", false
}

// 获取槽位的值，默认取第一组槽位
func (this *Dialog) GetSlotValue(name string, index ...int) string {
	// 默认取第一个intent
	// 如果有第二个参数，取参数指定index
	// 如果有超过三个参数，从第三个参数开始忽略
	i := getIndex(index)

	intent := this.getIntent(i)
	if intent != nil {
		return intent.GetSlotValue(name)
	}
	return ""
}

// 获取槽位的确认状态，默认取第一组槽位
func (this *Dialog) GetSlotConfirmationStatus(name string, index ...int) string {
	i := getIndex(index)
	intent := this.getIntent(i)

	if intent != nil {
		return intent.GetSlotStatus(name)
	}
	return ""
}

// 获取意图的确认状态
func (this *Dialog) GetIntentConfirmationStatus(index ...int) string {
	i := getIndex(index)
	intent := this.getIntent(i)

	if intent != nil {
		return intent.ConfirmationStatus
	}
	return ""
}

// 托管对话. 对话由DuerOS代为处理
func (this *Dialog) Delegate() *Dialog {
	this.directive = &data.DialogDirective{
		Type:          "Dialog.Delegate",
		UpdatedIntent: this.getIntent(0).GetData(),
	}
	return this
}

func (this *Dialog) ConfirmSlot(name string) *Dialog {
	intent := this.getIntent(0)
	slot := intent.GetSlot(name)

	if slot != nil {
		this.directive = &data.DialogDirective{
			Type:          "Dialog.ConfirmSlot",
			SlotToConfirm: name,
			UpdatedIntent: intent.GetData(),
		}
	}
	return this
}

func (this *Dialog) ConfirmIntent() *Dialog {
	this.directive = &data.DialogDirective{
		Type:          "Dialog.ConfirmIntent",
		UpdatedIntent: this.getIntent(0).GetData(),
	}
	return this
}

func (this *Dialog) ElicitSlot(name string) *Dialog {
	intent := this.getIntent(0)
	slot := intent.GetSlot(name)

	if slot != nil {
		this.directive = &data.DialogDirective{
			Type:          "Dialog.ElicitSlot",
			SlotToElicit:  name,
			UpdatedIntent: intent.GetData(),
		}
	}
	return this
}

func (this *Dialog) GetDirective() *data.DialogDirective {
	return this.directive
}
