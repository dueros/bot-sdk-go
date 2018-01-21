package nlu

import (
	"../data"
)

type Dialog struct {
	data        data.IntentRequestBody
	Intents     []*Intent
	DialogState string
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

func (this *Dialog) GetQuery() (string, bool) {
	if this.data.Query.Type == "TEXT" {
		return this.data.Query.Original, true
	}
	return "", false
}

func (this *Dialog) GetIntentName() (string, bool) {
	l := len(this.Intents)
	if l > 0 {
		intent := this.Intents[0]
		return intent.Name, true
	}
	return "", false
}
