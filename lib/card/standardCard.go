package card

type StandardCard struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
	BaseCard
}

func NewStandardCard() *StandardCard {
	card := &StandardCard{}
	card.Type = "standard"
	return card
}

func (this *StandardCard) SetTitle(title string) *StandardCard {
	this.Title = title
	return this
}

func (this *StandardCard) SetContent(content string) *StandardCard {
	this.Content = content
	return this
}

func (this *StandardCard) SetImage(image string) *StandardCard {
	this.Image = image
	return this
}
