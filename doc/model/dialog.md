

# model
`import "github.com/dueros/bot-sdk-go/bot/model"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type Dialog](#Dialog)
  * [func NewDialog(request data.IntentRequestBody) *Dialog](#NewDialog)
  * [func (this *Dialog) ConfirmIntent() *Dialog](#Dialog.ConfirmIntent)
  * [func (this *Dialog) ConfirmSlot(name string) *Dialog](#Dialog.ConfirmSlot)
  * [func (this *Dialog) Delegate() *Dialog](#Dialog.Delegate)
  * [func (this *Dialog) ElicitSlot(name string) *Dialog](#Dialog.ElicitSlot)
  * [func (this *Dialog) GetDirective() *data.DialogDirective](#Dialog.GetDirective)
  * [func (this *Dialog) GetIntentConfirmationStatus(index ...int) string](#Dialog.GetIntentConfirmationStatus)
  * [func (this *Dialog) GetIntentName() (string, bool)](#Dialog.GetIntentName)
  * [func (this *Dialog) GetQuery() (string, bool)](#Dialog.GetQuery)
  * [func (this *Dialog) GetSlotConfirmationStatus(name string, index ...int) string](#Dialog.GetSlotConfirmationStatus)
  * [func (this *Dialog) GetSlotValue(name string, index ...int) string](#Dialog.GetSlotValue)
* [type IntentRequest](#IntentRequest)
  * [func (this *IntentRequest) GetIntentName() (string, bool)](#IntentRequest.GetIntentName)
  * [func (this *IntentRequest) GetQuery() string](#IntentRequest.GetQuery)
  * [func (this *IntentRequest) IsDialogStateCompleted() bool](#IntentRequest.IsDialogStateCompleted)


#### <a name="pkg-files">Package files</a>
[dialog.go](/src/github.com/dueros/bot-sdk-go/bot/model/dialog.go) [intent.go](/src/github.com/dueros/bot-sdk-go/bot/model/intent.go) [request.go](/src/github.com/dueros/bot-sdk-go/bot/model/request.go) [response.go](/src/github.com/dueros/bot-sdk-go/bot/model/response.go) [session.go](/src/github.com/dueros/bot-sdk-go/bot/model/session.go) [ssml_builder.go](/src/github.com/dueros/bot-sdk-go/bot/model/ssml_builder.go) 






## <a name="Dialog">type</a> [Dialog](/src/target/dialog.go?s=68:204#L7)
``` go
type Dialog struct {
    Intents     []*Intent
    DialogState string
    // contains filtered or unexported fields
}
```






### <a name="NewDialog">func</a> [NewDialog](/src/target/dialog.go?s=206:260#L14)
``` go
func NewDialog(request data.IntentRequestBody) *Dialog
```




### <a name="Dialog.ConfirmIntent">func</a> (\*Dialog) [ConfirmIntent](/src/target/dialog.go?s=2601:2644#L122)
``` go
func (this *Dialog) ConfirmIntent() *Dialog
```
对意图进行确认




### <a name="Dialog.ConfirmSlot">func</a> (\*Dialog) [ConfirmSlot](/src/target/dialog.go?s=2278:2330#L107)
``` go
func (this *Dialog) ConfirmSlot(name string) *Dialog
```
对槽位进行confirm




### <a name="Dialog.Delegate">func</a> (\*Dialog) [Delegate](/src/target/dialog.go?s=2069:2107#L98)
``` go
func (this *Dialog) Delegate() *Dialog
```
托管对话. 对话由DuerOS代为处理




### <a name="Dialog.ElicitSlot">func</a> (\*Dialog) [ElicitSlot](/src/target/dialog.go?s=2810:2861#L131)
``` go
func (this *Dialog) ElicitSlot(name string) *Dialog
```
询问槽位




### <a name="Dialog.GetDirective">func</a> (\*Dialog) [GetDirective](/src/target/dialog.go?s=3107:3163#L145)
``` go
func (this *Dialog) GetDirective() *data.DialogDirective
```



### <a name="Dialog.GetIntentConfirmationStatus">func</a> (\*Dialog) [GetIntentConfirmationStatus](/src/target/dialog.go?s=1829:1897#L87)
``` go
func (this *Dialog) GetIntentConfirmationStatus(index ...int) string
```
获取意图的确认状态




### <a name="Dialog.GetIntentName">func</a> (\*Dialog) [GetIntentName](/src/target/dialog.go?s=992:1042#L53)
``` go
func (this *Dialog) GetIntentName() (string, bool)
```
获取当前意图的名字




### <a name="Dialog.GetQuery">func</a> (\*Dialog) [GetQuery](/src/target/dialog.go?s=812:857#L45)
``` go
func (this *Dialog) GetQuery() (string, bool)
```
获取用户请求的原始query




### <a name="Dialog.GetSlotConfirmationStatus">func</a> (\*Dialog) [GetSlotConfirmationStatus](/src/target/dialog.go?s=1591:1670#L76)
``` go
func (this *Dialog) GetSlotConfirmationStatus(name string, index ...int) string
```
获取槽位的确认状态，默认取第一组槽位




### <a name="Dialog.GetSlotValue">func</a> (\*Dialog) [GetSlotValue](/src/target/dialog.go?s=1194:1260#L62)
``` go
func (this *Dialog) GetSlotValue(name string, index ...int) string
```
获取槽位的值，默认取第一组槽位




## <a name="IntentRequest">type</a> [IntentRequest](/src/target/request.go?s=1244:1325#L35)
``` go
type IntentRequest struct {
    Data   data.IntentRequest
    Dialog *Dialog
    Request
}
```









### <a name="IntentRequest.GetIntentName">func</a> (\*IntentRequest) [GetIntentName](/src/target/request.go?s=1551:1608#L57)
``` go
func (this *IntentRequest) GetIntentName() (string, bool)
```
获取意图名




### <a name="IntentRequest.GetQuery">func</a> (\*IntentRequest) [GetQuery](/src/target/request.go?s=1780:1824#L67)
``` go
func (this *IntentRequest) GetQuery() string
```
获取用户请求query




### <a name="IntentRequest.IsDialogStateCompleted">func</a> (\*IntentRequest) [IsDialogStateCompleted](/src/target/request.go?s=1678:1734#L62)
``` go
func (this *IntentRequest) IsDialogStateCompleted() bool
```
槽位填充是否完成








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
