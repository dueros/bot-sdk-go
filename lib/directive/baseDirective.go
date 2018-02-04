package directive

type BaseDirective struct {
	Type string `json:"type"`
}

func (this *BaseDirective) GenToken() string {
	return "uuid"
}
