package data

type Session struct {
	New        bool
	SessionId  string
	Attributes map[string]string
}

type System struct {
	User        User
	Application Application
	Device      Device
}

type User struct {
	UserId   string
	UserInfo UserInfo
}

type UserInfo struct {
}

type Application struct {
	ApplicationId string
}

type Device struct {
	DeviceId string
}

type Context struct {
	System System
}

type baseRequest struct {
	Type      string
	RequestId string
	Timestamp string
}

type RequestPart struct {
	Version string
	Session Session
	Context Context
}

// 打开请求
type LaunchRequest struct {
	RequestPart
	Request baseRequest
}

// 退出请求
type SessionEndedRequestBody struct {
	baseRequest
	Reason string
}

type SessionEndedRequest struct {
	RequestPart
	Request SessionEndedRequestBody
}

// intent request
type Query struct {
	Type     string
	Original string
}

type Slot struct {
	Name               string
	Value              string
	ConfirmationStatus string
}

type Intent struct {
	Name               string
	Slots              map[string]Slot
	ConfirmationStatus string
}

type IntentRequestBody struct {
	baseRequest
	Query       Query
	DialogState string
	Intents     []Intent
}

type IntentRequest struct {
	RequestPart
	Request IntentRequestBody
}
