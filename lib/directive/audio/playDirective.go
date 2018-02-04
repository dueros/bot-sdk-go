package audio

import (
	"../../directive"
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
