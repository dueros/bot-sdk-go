package template

type ListItem struct {
	Token       string `json:"token"`
	Image       *Image `json:"image,omitempty"`
	TextContent struct {
		PrimaryText   *Text `json:"primaryText"`
		SecondaryText *Text `json:"secondaryText"`
		TertiaryText  *Text `json:"tertiaryText,omitempty"`
	} `json:"textContent"`
}

func NewListItem() *ListItem {
	item := &ListItem{}
	return item
}

func (this *ListItem) SetImage(image *Image) *ListItem {
	this.Image = image
	return this
}

func (this *ListItem) SetImageUrl(url string) *ListItem {
	image := NewImage(url)
	return this.SetImage(image)
}

func (this *ListItem) SetPlainPrimaryText(text string) *ListItem {
	this.TextContent.PrimaryText = NewText(PLAIN_TEXT, text)
	return this
}

func (this *ListItem) SetPlainSecondaryText(text string) *ListItem {
	this.TextContent.SecondaryText = NewText(PLAIN_TEXT, text)
	return this
}

func (this *ListItem) SetPlainTertiary(text string) *ListItem {
	this.TextContent.TertiaryText = NewText(PLAIN_TEXT, text)
	return this
}

type ListTemplate struct {
	BaseTemplate
	ListItems []*ListItem `json:"listItems"`
}

func (this *ListTemplate) AddItem(listItem *ListItem) {
	this.ListItems = append(this.ListItems, listItem)
}
