package card

type ImageCard struct {
	BaseCard
	list []map[string]string
}

func NewImageCard() *ImageCard {
	card := &ImageCard{}
	card.Type = "image"
	return card
}

func (this *ImageCard) AddItem(src string, thumbnail string) *ImageCard {
	this.list = append(this.list, map[string]string{
		"src":       src,
		"thumbnail": thumbnail,
	})
	return this
}
