package model

import (
	"github.com/dueros/bot-sdk-go/bot/data"
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

// 当前session是否是新的
func (this *Session) IsNew() bool {
	return this.data.New
}

// 获取session id
func (this *Session) GetId() string {
	return this.data.SessionId
}

// 获取session中对应字段的值
func (this *Session) GetAttribute(key string) string {
	value, ok := this.data.Attributes[key]
	if ok {
		return value
	}
	return ""
}

// 设置session中对应字段的值
func (this *Session) SetAttribute(key, value string) {
	if key == "" {
		return
	}

	this.data.Attributes[key] = value
}

func (this *Session) GetData() data.Session {
	return this.data
}
