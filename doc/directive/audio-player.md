

# audio_player
`import "github.com/dueros/bot-sdk-go/bot/directive/audio_player"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type PlayDirective](#PlayDirective)
  * [func NewPlayDirective(url string) *PlayDirective](#NewPlayDirective)
  * [func (this *PlayDirective) GetToken(token string) string](#PlayDirective.GetToken)
  * [func (this *PlayDirective) SetBehavior(behavior string) *PlayDirective](#PlayDirective.SetBehavior)
  * [func (this *PlayDirective) SetOffsetInMilliseconds(milliseconds int) *PlayDirective](#PlayDirective.SetOffsetInMilliseconds)
  * [func (this *PlayDirective) SetPlayerInfo(playerInfo *PlayerInfo) *PlayDirective](#PlayDirective.SetPlayerInfo)
  * [func (this *PlayDirective) SetProgressReportIntervalMs(intervalMs int) *PlayDirective](#PlayDirective.SetProgressReportIntervalMs)
  * [func (this *PlayDirective) SetStreamFormat(format string) *PlayDirective](#PlayDirective.SetStreamFormat)
  * [func (this *PlayDirective) SetToken(token string) *PlayDirective](#PlayDirective.SetToken)
  * [func (this *PlayDirective) SetUrl(url string) *PlayDirective](#PlayDirective.SetUrl)
* [type PlayerInfo](#PlayerInfo)
  * [func NewPlayerInfo() *PlayerInfo](#NewPlayerInfo)
  * [func (this *PlayerInfo) SetArt(src string) *PlayerInfo](#PlayerInfo.SetArt)
  * [func (this *PlayerInfo) SetLyric(url string) *PlayerInfo](#PlayerInfo.SetLyric)
  * [func (this *PlayerInfo) SetMediaLengthInMs(length uint64) *PlayerInfo](#PlayerInfo.SetMediaLengthInMs)
  * [func (this *PlayerInfo) SetProviderLogo(logo string) *PlayerInfo](#PlayerInfo.SetProviderLogo)
  * [func (this *PlayerInfo) SetProviderName(name string) *PlayerInfo](#PlayerInfo.SetProviderName)
  * [func (this *PlayerInfo) SetTitle(title string) *PlayerInfo](#PlayerInfo.SetTitle)
  * [func (this *PlayerInfo) SetTitleSubtext1(text string) *PlayerInfo](#PlayerInfo.SetTitleSubtext1)
  * [func (this *PlayerInfo) SetTitleSubtext2(text string) *PlayerInfo](#PlayerInfo.SetTitleSubtext2)
* [type StopDirective](#StopDirective)
  * [func NewStopDirective() *StopDirective](#NewStopDirective)


#### <a name="pkg-files">Package files</a>
[const.go](/src/github.com/dueros/bot-sdk-go/bot/directive/audio_player/const.go) [play.go](/src/github.com/dueros/bot-sdk-go/bot/directive/audio_player/play.go) [player_info.go](/src/github.com/dueros/bot-sdk-go/bot/directive/audio_player/player_info.go) [stop.go](/src/github.com/dueros/bot-sdk-go/bot/directive/audio_player/stop.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    REPLACE_ALL      = "REPLACE_ALL"
    REPLACE_ENQUEUED = "REPLACE_ENQUEUED"
    ENQUEUE          = "ENQUEUE"

    AUDIO_MP3  = "AUDIO_MP3"
    AUDIO_M3U8 = "AUDIO_M3U8"
    AUDIO_M4A  = "AUDIO_M4A"

    AUDIO_TYPE_MUSIC = "MUSIC"

    LYRIC_FORMAT_LRC = "LRC"
)
```




## <a name="PlayDirective">type</a> [PlayDirective](/src/target/play.go?s=291:825#L19)
``` go
type PlayDirective struct {
    directive.BaseDirective
    PlayBehavior string `json:"playBehavior"`
    AudioItem    struct {
        Stream struct {
            Url                      string `json:"url"`
            StreamFormat             string `json:"streamFormat"`
            OffsetInMilliseconds     int    `json:"offsetInMilliSeconds"`
            Token                    string `json:"token"`
            ProgressReportIntervalMs int    `json:"progressReportIntervalMs,omitempty"`
        } `json:"stream"`
        PlayerInfo *PlayerInfo `json:"playerInfo,omitempty"`
    } `json:"audioItem"`
}
```






### <a name="NewPlayDirective">func</a> [NewPlayDirective](/src/target/play.go?s=827:875#L34)
``` go
func NewPlayDirective(url string) *PlayDirective
```




### <a name="PlayDirective.GetToken">func</a> (\*PlayDirective) [GetToken](/src/target/play.go?s=1534:1590#L61)
``` go
func (this *PlayDirective) GetToken(token string) string
```



### <a name="PlayDirective.SetBehavior">func</a> (\*PlayDirective) [SetBehavior](/src/target/play.go?s=1249:1319#L47)
``` go
func (this *PlayDirective) SetBehavior(behavior string) *PlayDirective
```



### <a name="PlayDirective.SetOffsetInMilliseconds">func</a> (\*PlayDirective) [SetOffsetInMilliseconds](/src/target/play.go?s=1744:1827#L70)
``` go
func (this *PlayDirective) SetOffsetInMilliseconds(milliseconds int) *PlayDirective
```



### <a name="PlayDirective.SetPlayerInfo">func</a> (\*PlayDirective) [SetPlayerInfo](/src/target/play.go?s=2252:2331#L88)
``` go
func (this *PlayDirective) SetPlayerInfo(playerInfo *PlayerInfo) *PlayDirective
```



### <a name="PlayDirective.SetProgressReportIntervalMs">func</a> (\*PlayDirective) [SetProgressReportIntervalMs](/src/target/play.go?s=1905:1990#L75)
``` go
func (this *PlayDirective) SetProgressReportIntervalMs(intervalMs int) *PlayDirective
```



### <a name="PlayDirective.SetStreamFormat">func</a> (\*PlayDirective) [SetStreamFormat](/src/target/play.go?s=2070:2142#L80)
``` go
func (this *PlayDirective) SetStreamFormat(format string) *PlayDirective
```



### <a name="PlayDirective.SetToken">func</a> (\*PlayDirective) [SetToken](/src/target/play.go?s=1414:1478#L56)
``` go
func (this *PlayDirective) SetToken(token string) *PlayDirective
```



### <a name="PlayDirective.SetUrl">func</a> (\*PlayDirective) [SetUrl](/src/target/play.go?s=1632:1692#L65)
``` go
func (this *PlayDirective) SetUrl(url string) *PlayDirective
```



## <a name="PlayerInfo">type</a> [PlayerInfo](/src/target/player_info.go?s=22:665#L3)
``` go
type PlayerInfo struct {
    Content struct {
        AudioItemType string `json:"audioItemType"`
        Title         string `json:"title"`
        TitleSubtext1 string `json:"titleSubtext1"`
        TitleSubtext2 string `json:"titleSubtext2"`
        Lyric         struct {
            Url    string `json:"url"`
            Format string `json:"format"`
        } `json:"lyric"`
        MediaLengthInMilliseconds uint64 `json:"mediaLengthInMilliseconds,omitempty"`
        Art                       struct {
            Src string `json:"src"`
        } `json:"art"`
        Provider struct {
            Name string `json:"name"`
            Logo struct {
                Src string `json:"src"`
            } `json:"logo"`
        } `json:"provider"`
    } `json:"content"`
}
```






### <a name="NewPlayerInfo">func</a> [NewPlayerInfo](/src/target/player_info.go?s=667:699#L26)
``` go
func NewPlayerInfo() *PlayerInfo
```




### <a name="PlayerInfo.SetArt">func</a> (\*PlayerInfo) [SetArt](/src/target/player_info.go?s=1439:1493#L60)
``` go
func (this *PlayerInfo) SetArt(src string) *PlayerInfo
```



### <a name="PlayerInfo.SetLyric">func</a> (\*PlayerInfo) [SetLyric](/src/target/player_info.go?s=1151:1207#L49)
``` go
func (this *PlayerInfo) SetLyric(url string) *PlayerInfo
```



### <a name="PlayerInfo.SetMediaLengthInMs">func</a> (\*PlayerInfo) [SetMediaLengthInMs](/src/target/player_info.go?s=1302:1371#L55)
``` go
func (this *PlayerInfo) SetMediaLengthInMs(length uint64) *PlayerInfo
```



### <a name="PlayerInfo.SetProviderLogo">func</a> (\*PlayerInfo) [SetProviderLogo](/src/target/player_info.go?s=1658:1722#L70)
``` go
func (this *PlayerInfo) SetProviderLogo(logo string) *PlayerInfo
```



### <a name="PlayerInfo.SetProviderName">func</a> (\*PlayerInfo) [SetProviderName](/src/target/player_info.go?s=1540:1604#L65)
``` go
func (this *PlayerInfo) SetProviderName(name string) *PlayerInfo
```



### <a name="PlayerInfo.SetTitle">func</a> (\*PlayerInfo) [SetTitle](/src/target/player_info.go?s=808:866#L34)
``` go
func (this *PlayerInfo) SetTitle(title string) *PlayerInfo
```



### <a name="PlayerInfo.SetTitleSubtext1">func</a> (\*PlayerInfo) [SetTitleSubtext1](/src/target/player_info.go?s=913:978#L39)
``` go
func (this *PlayerInfo) SetTitleSubtext1(text string) *PlayerInfo
```



### <a name="PlayerInfo.SetTitleSubtext2">func</a> (\*PlayerInfo) [SetTitleSubtext2](/src/target/player_info.go?s=1032:1097#L44)
``` go
func (this *PlayerInfo) SetTitleSubtext2(text string) *PlayerInfo
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
