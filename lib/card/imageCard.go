package card

type ImageCard struct {
	BaseCard
	List []map[string]string `json:"list"`
}

func NewImageCard() *ImageCard {
	card := &ImageCard{}
	card.Type = "image"
	return card
}

func (this *ImageCard) AddItem(src string, thumbnail string) *ImageCard {
	this.List = append(this.List, map[string]string{
		"src":       src,
		"thumbnail": thumbnail,
	})
	return this
}
