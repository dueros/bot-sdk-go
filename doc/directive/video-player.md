

# video_player
`import "github.com/dueros/bot-sdk-go/bot/directive/video_player"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type PlayDirective](#PlayDirective)
  * [func NewPlayDirective(url string) *PlayDirective](#NewPlayDirective)
  * [func (this *PlayDirective) GetToken(token string) string](#PlayDirective.GetToken)
  * [func (this *PlayDirective) SetBehavior(behavior string) *PlayDirective](#PlayDirective.SetBehavior)
  * [func (this *PlayDirective) SetExpectedPreviousToken(expectedPreviousToken string) *PlayDirective](#PlayDirective.SetExpectedPreviousToken)
  * [func (this *PlayDirective) SetExpiryTime(expiryTime string) *PlayDirective](#PlayDirective.SetExpiryTime)
  * [func (this *PlayDirective) SetOffsetInMilliseconds(milliseconds int) *PlayDirective](#PlayDirective.SetOffsetInMilliseconds)
  * [func (this *PlayDirective) SetReportDelayInMs(reportDelayInMs int) *PlayDirective](#PlayDirective.SetReportDelayInMs)
  * [func (this *PlayDirective) SetReportIntervalInMs(reportIntervalInMs int) *PlayDirective](#PlayDirective.SetReportIntervalInMs)
  * [func (this *PlayDirective) SetToken(token string) *PlayDirective](#PlayDirective.SetToken)
  * [func (this *PlayDirective) SetUrl(url string) *PlayDirective](#PlayDirective.SetUrl)
* [type StopDirective](#StopDirective)
  * [func NewStopDirective() *StopDirective](#NewStopDirective)


#### <a name="pkg-files">Package files</a>
[const.go](/src/github.com/dueros/bot-sdk-go/bot/directive/video_player/const.go) [play.go](/src/github.com/dueros/bot-sdk-go/bot/directive/video_player/play.go) [stop.go](/src/github.com/dueros/bot-sdk-go/bot/directive/video_player/stop.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    REPLACE_ALL      = "REPLACE_ALL"
    REPLACE_ENQUEUED = "REPLACE_ENQUEUED"
    ENQUEUE          = "ENQUEUE"
)
```




## <a name="PlayDirective">type</a> [PlayDirective](/src/target/play.go?s=193:930#L13)
``` go
type PlayDirective struct {
    directive.BaseDirective
    PlayBehavior string `json:"playBehavior"`
    VideoItem    struct {
        Stream struct {
            Url                  string `json:"url"`
            OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
            ExpiryTime           string `json:"expiryTime,omitempty"`
            ProgressReport       struct {
                ProgressReportDelayInMilliseconds    int `json:"progressReportDelayInMilliseconds,omitempty"`
                ProgressReportIntervalInMilliseconds int `json:"progressReportIntervalInMilliseconds,omitempty"`
            } `json:"progressReport,omitempty"`
            Token                 string `json:"token"`
            ExpectedPreviousToken string `json:"expectedPreviousToken,omitempty"`
        } `json:"stream"`
    } `json:"VideoItem"`
}
```






### <a name="NewPlayDirective">func</a> [NewPlayDirective](/src/target/play.go?s=932:980#L31)
``` go
func NewPlayDirective(url string) *PlayDirective
```




### <a name="PlayDirective.GetToken">func</a> (\*PlayDirective) [GetToken](/src/target/play.go?s=1503:1559#L55)
``` go
func (this *PlayDirective) GetToken(token string) string
```



### <a name="PlayDirective.SetBehavior">func</a> (\*PlayDirective) [SetBehavior](/src/target/play.go?s=1218:1288#L41)
``` go
func (this *PlayDirective) SetBehavior(behavior string) *PlayDirective
```



### <a name="PlayDirective.SetExpectedPreviousToken">func</a> (\*PlayDirective) [SetExpectedPreviousToken](/src/target/play.go?s=2406:2502#L84)
``` go
func (this *PlayDirective) SetExpectedPreviousToken(expectedPreviousToken string) *PlayDirective
```



### <a name="PlayDirective.SetExpiryTime">func</a> (\*PlayDirective) [SetExpiryTime](/src/target/play.go?s=1874:1948#L69)
``` go
func (this *PlayDirective) SetExpiryTime(expiryTime string) *PlayDirective
```



### <a name="PlayDirective.SetOffsetInMilliseconds">func</a> (\*PlayDirective) [SetOffsetInMilliseconds](/src/target/play.go?s=1713:1796#L64)
``` go
func (this *PlayDirective) SetOffsetInMilliseconds(milliseconds int) *PlayDirective
```



### <a name="PlayDirective.SetReportDelayInMs">func</a> (\*PlayDirective) [SetReportDelayInMs](/src/target/play.go?s=2014:2095#L74)
``` go
func (this *PlayDirective) SetReportDelayInMs(reportDelayInMs int) *PlayDirective
```



### <a name="PlayDirective.SetReportIntervalInMs">func</a> (\*PlayDirective) [SetReportIntervalInMs](/src/target/play.go?s=2204:2291#L79)
``` go
func (this *PlayDirective) SetReportIntervalInMs(reportIntervalInMs int) *PlayDirective
```



### <a name="PlayDirective.SetToken">func</a> (\*PlayDirective) [SetToken](/src/target/play.go?s=1383:1447#L50)
``` go
func (this *PlayDirective) SetToken(token string) *PlayDirective
```



### <a name="PlayDirective.SetUrl">func</a> (\*PlayDirective) [SetUrl](/src/target/play.go?s=1601:1661#L59)
``` go
func (this *PlayDirective) SetUrl(url string) *PlayDirective
```



## <a name="StopDirective">type</a> [StopDirective](/src/target/stop.go?s=80:134#L7)
``` go
type StopDirective struct {
    directive.BaseDirective
}
```






### <a name="NewStopDirective">func</a> [NewStopDirective](/src/target/stop.go?s=136:174#L11)
``` go
func NewStopDirective() *StopDirective
```








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
