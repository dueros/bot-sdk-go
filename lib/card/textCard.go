package card

type TextCard struct {
	BaseCard
	content string `json:"content"`
}

func NewTextCard(content string) *TextCard {
	card := &TextCard{}
	card.Type = "txt"
	card.SetContent(content)
	return card
}

func (this *TextCard) SetContent(content string) {
	this.content = content
}
