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
	UserId      string
	AccessToken string
	UserInfo    UserInfo
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
	System      System
	AudioPlayer AudioPlayerContext
}

type AudioPlayerContext struct {
	Token                string
	OffsetInMilliSeconds int32
	PlayerActivity       string
}

type baseRequest struct {
	Type      string
	RequestId string
	Timestamp string
}

// 公共请求体
type RequestPart struct {
	Version string
	Session Session
	Context Context
	Request baseRequest
}

// 事件请求
type EventRequest struct {
	Request struct {
		baseRequest
		Token                string
		OffsetInMilliSeconds int32 `json:"offsetInMilliSeconds,omitempty"` //Audio Player Event
	}
}

// 打开请求
type LaunchRequest struct {
	Request baseRequest
}

// 退出请求
type SessionEndedRequestBody struct {
	baseRequest
	Reason string
}

type SessionEndedRequest struct {
	Request SessionEndedRequestBody
}

// intent request
type Query struct {
	Type     string
	Original string
}

type Slot struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	ConfirmationStatus string `json:"confirmationStatus"`
}

type Intent struct {
	Name               string          `json:"name"`
	Slots              map[string]Slot `json:"slots"`
	ConfirmationStatus string          `json:"confirmationStatus"`
}

type IntentRequestBody struct {
	baseRequest
	Query       Query
	DialogState string
	Intents     []Intent
}

type IntentRequest struct {
	Request IntentRequestBody
}
