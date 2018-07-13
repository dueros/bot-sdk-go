package util

import (
	"regexp"

	"github.com/dueros/bot-sdk-go/bot/data"
)

func FormatSpeech(speech string) data.Speech {
	match, _ := regexp.MatchString("^<speak>", speech)

	if match {
		return data.Speech{
			Type: "SSML",
			Ssml: speech,
		}
	}

	return data.Speech{
		Type: "PlainText",
		Text: speech,
	}
}
