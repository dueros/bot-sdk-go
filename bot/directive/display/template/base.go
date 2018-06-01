package template

import (
	"github.com/dueros/bot-sdk-go/bot/directive"
)

const (
	PLAIN_TEXT = "PlainText"
	RICH_TEXT  = "RichText"
)

var textTypeMap = map[string]bool{
	PLAIN_TEXT: true,
	RICH_TEXT:  true,
}

type BaseTemplate struct {
	directive.BaseDirective
	Token           string `json:"token"`
	BackgroundImage *Image `json:"backgroundImage,omitempty"`
	Title           string `json:"title"`
}

func (this *BaseTemplate) SetTitle(title string) {
	this.Title = title
}

func (this *BaseTemplate) SetBackgroundImageUrl(url string) {
	this.SetBackgroundImage(NewImage(url))
}

func (this *BaseTemplate) SetBackgroundImage(background *Image) {
	this.BackgroundImage = background
}

type Image struct {
	Url          string `json:"url"`
	WidthPixels  int    `json:"widthPixels,omitempty"`
	HeightPixels int    `json:"heightPixels,omitempty"`
}

func NewImage(url string) *Image {
	image := &Image{}
	image.Url = url
	return image
}

func (this *Image) SetWidth(width int) *Image {
	this.WidthPixels = width
	return this
}

func (this *Image) SetHeight(height int) *Image {
	this.HeightPixels = height
	return this
}

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewText(textType, text string) *Text {
	ok := textTypeMap[textType]
	if !ok {
		textType = PLAIN_TEXT
	}

	t := &Text{}
	t.Type = textType
	t.Text = text
	return t
}
