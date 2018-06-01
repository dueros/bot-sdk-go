package template

type BodyTemplate5 struct {
	BaseTemplate
	Images []*Image `json:"images"`
}

func NewBodyTemplate5() *BodyTemplate5 {
	bodyTemplate := &BodyTemplate5{}
	bodyTemplate.Images = make([]*Image, 0)
	return bodyTemplate
}

func (this *BodyTemplate5) AddImage(image *Image) *BodyTemplate5 {
	this.Images = append(this.Images, image)
	return this
}
