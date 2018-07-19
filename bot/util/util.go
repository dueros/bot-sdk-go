package util

import (
	"io/ioutil"
	"os"
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

func ReadFileAll(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}
