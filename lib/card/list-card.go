package card

type ListCardItem struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"content"`
	Url     string `json:"url"`
}

func NewListCardItem() *ListCardItem {
	card := &ListCardItem{}
	return card
}

func (this *ListCardItem) SetTitle(title string) *ListCardItem {
	this.Title = title
	return this
}

func (this *ListCardItem) SetContent(content string) *ListCardItem {
	this.Content = content
	return this
}

func (this *ListCardItem) SetImage(image string) *ListCardItem {
	this.Image = image
	return this
}

func (this *ListCardItem) SetUrl(url string) *ListCardItem {
	this.Url = url
	return this
}

type ListCard struct {
	List []*ListCardItem `json:"list"`
	BaseCard
}

func NewListCard() *ListCard {
	card := &ListCard{}
	card.Type = "list"
	return card
}

func (this *ListCard) AddItem(item *ListCardItem) *ListCard {
	this.List = append(this.List, item)
	return this
}
