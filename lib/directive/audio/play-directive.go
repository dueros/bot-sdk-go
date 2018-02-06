package audio

import (
	"github.com/dueros/bot-sdk-go/lib/directive"
)

type PlayDirective struct {
	directive.BaseDirective
	PlayBehavior string `json:"playBehavior"`
}

func NewPlayDirective() *PlayDirective {
	directive := &PlayDirective{}
	directive.Type = "tst"
	directive.PlayBehavior = "ga"
	return directive
}
