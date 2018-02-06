package card

type TextCard struct {
	BaseCard
	Content string `json:"content"`
}

func NewTextCard(content string) *TextCard {
	card := &TextCard{}
	card.Type = "txt"
	card.SetContent(content)
	return card
}

func (this *TextCard) SetContent(content string) {
	this.Content = content
}
