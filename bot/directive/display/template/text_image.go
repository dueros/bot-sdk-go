package template

type TextImageTemplate struct {
	BaseTemplate
	Content *Text  `json:"content"`
	Image   *Image `json:"image"`
}

func (this *TextImageTemplate) SetPlainContent(content string) {
	text := NewText(PLAIN_TEXT, content)
	this.Content = text
}

func (this *TextImageTemplate) SetImageUrl(url string) {
	image := NewImage(url)
	this.SetImage(image)
}

func (this *TextImageTemplate) SetImage(image *Image) {
	this.Image = image
}
