package card

type ListCardItem struct {
	title   string `json:"title"`
	content string `json:"content"`
	image   string `json:"content"`
	url     string `json:"url"`
}

func NewListCardItem() *ListCardItem {
	card := &ListCardItem{}
	return card
}

func (this *ListCardItem) SetTitle(title string) *ListCardItem {
	this.title = title
	return this
}

func (this *ListCardItem) SetContent(content string) *ListCardItem {
	this.content = content
	return this
}

func (this *ListCardItem) SetImage(image string) *ListCardItem {
	this.image = image
	return this
}

func (this *ListCardItem) SetUrl(url string) *ListCardItem {
	this.url = url
	return this
}

type ListCard struct {
	list []*ListCardItem
	BaseCard
}

func NewListCard() *ListCard {
	card := &ListCard{}
	card.Type = "list"
	return card
}

func (this *ListCard) AddItem(item *ListCardItem) *ListCard {
	this.list = append(this.list, item)
	return this
}
