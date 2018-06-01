package template

const (
	BOTTOM_LEFT_POSITION = "BOTTOM-LEFT"
	CENTER               = "CENTER"
	TOP_LEFT             = "TOP-LEFT"
)

var textPositionMap = map[string]bool{
	BOTTOM_LEFT_POSITION: true,
	CENTER:               true,
	TOP_LEFT:             true,
}

type BodyTemplate1 struct {
	BaseTemplate
	TextContent struct {
		Position string `json:"position"`
		Text     *Text  `json:"text"`
	} `json:"content"`
}

func NewBodyTemplate1() *BodyTemplate1 {
	bodyTemplate := &BodyTemplate1{}
	bodyTemplate.Type = "BodyTemplate1"
	bodyTemplate.Token = bodyTemplate.GenToken()
	bodyTemplate.TextContent.Position = BOTTOM_LEFT_POSITION
	return bodyTemplate
}

func (this *BodyTemplate1) SetContentPosition(position string) *BodyTemplate1 {
	ok := textPositionMap[position]
	if !ok {
		position = BOTTOM_LEFT_POSITION
	}

	this.TextContent.Position = position
	return this
}

func (this *BodyTemplate1) SetPlainContent(content string) *BodyTemplate1 {
	text := NewText(PLAIN_TEXT, content)
	this.TextContent.Text = text
	return this
}
