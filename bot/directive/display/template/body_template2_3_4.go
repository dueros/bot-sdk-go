package template

type BodyTemplate2 struct {
	TextImageTemplate
}

func NewBodyTemplate2() *BodyTemplate2 {
	bodyTemplate := &BodyTemplate2{}
	bodyTemplate.Type = "BodyTemplate2"
	bodyTemplate.Token = bodyTemplate.GenToken()
	return bodyTemplate
}

type BodyTemplate3 struct {
	TextImageTemplate
}

func NewBodyTemplate3() *BodyTemplate3 {
	bodyTemplate := &BodyTemplate3{}
	bodyTemplate.Type = "BodyTemplate3"
	bodyTemplate.Token = bodyTemplate.GenToken()
	return bodyTemplate
}

type BodyTemplate4 struct {
	TextImageTemplate
}

func NewBodyTemplate4() *BodyTemplate4 {
	bodyTemplate := &BodyTemplate4{}
	bodyTemplate.Type = "BodyTemplate4"
	bodyTemplate.Token = bodyTemplate.GenToken()
	return bodyTemplate
}
