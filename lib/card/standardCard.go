package card

type StandardCard struct {
	title   string `json:"title"`
	content string `json:"content"`
	image   string `json:"image"`
	BaseCard
}

func NewStandardCard() *StandardCard {
	card := &StandardCard{}
	card.Type = "standard"
	return card
}

func (this *StandardCard) setTitle(title string) *StandardCard {
	this.title = title
	return this
}

func (this *StandardCard) setContent(content string) *StandardCard {
	this.content = content
	return this
}

func (this *StandardCard) setImage(image string) *StandardCard {
	this.image = image
	return this
}
