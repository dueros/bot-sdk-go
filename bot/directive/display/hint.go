package display

import (
	"github.com/dueros/bot-sdk-go/bot/data"
	"github.com/dueros/bot-sdk-go/bot/util"
)

type Hint struct {
	Type  string        `json:"type"`
	Hints []data.Speech `json:"hints"`
}

func NewHint(hint ...string) *Hint {
	h := &Hint{
		Type:  "Hint",
		Hints: make([]data.Speech, 0),
	}

	h.SetHints(hint)

	return h
}

func (this *Hint) SetHints(hints []string) *Hint {
	for _, value := range hints {
		this.AddHint(value)
	}

	return this
}

func (this *Hint) AddHint(hint string) *Hint {
	this.Hints = append(this.Hints, util.FormatSpeech(hint))

	return this
}
