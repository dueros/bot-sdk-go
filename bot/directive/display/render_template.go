package display

import (
	"github.com/dueros/bot-sdk-go/bot/directive"
)

type RenderTemplate struct {
	directive.BaseDirective
	Template interface{} `json:"template"`
}

func NewRenderTemplate(template interface{}) *RenderTemplate {
	renderTemplate := &RenderTemplate{}
	renderTemplate.Type = "Display.RenderTemplate"
	renderTemplate.Template = template
	return renderTemplate
}
