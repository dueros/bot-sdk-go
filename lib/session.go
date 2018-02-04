package lib

import (
	"./data"
)

type Session struct {
	data data.Session
}

func NewSession(data data.Session) *Session {
	if data.Attributes == nil {
		data.Attributes = make(map[string]string)
	}

	return &Session{
		data: data,
	}
}

func (this *Session) IsNew() bool {
	return this.data.New
}

func (this *Session) GetId() string {
	return this.data.SessionId
}

func (this *Session) GetAttribute(key string) string {
	value, ok := this.data.Attributes[key]
	if ok {
		return value
	}
	return ""
}

func (this *Session) SetAttribute(key, value string) {
	if key == "" {
		return
	}

	this.data.Attributes[key] = value
}

func (this *Session) GetData() data.Session {
	return this.data
}
