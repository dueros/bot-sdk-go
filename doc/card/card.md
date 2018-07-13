

# card
`import "github.com/dueros/bot-sdk-go/bot/card"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type BaseCard](#BaseCard)
* [type ImageCard](#ImageCard)
  * [func NewImageCard() *ImageCard](#NewImageCard)
  * [func (this *ImageCard) AddItem(src string, thumbnail string) *ImageCard](#ImageCard.AddItem)
* [type LinkAccountCard](#LinkAccountCard)
  * [func NewLinkAccountCard() *LinkAccountCard](#NewLinkAccountCard)
* [type ListCard](#ListCard)
  * [func NewListCard() *ListCard](#NewListCard)
  * [func (this *ListCard) AddItem(item *ListCardItem) *ListCard](#ListCard.AddItem)
* [type ListCardItem](#ListCardItem)
  * [func NewListCardItem() *ListCardItem](#NewListCardItem)
  * [func (this *ListCardItem) SetContent(content string) *ListCardItem](#ListCardItem.SetContent)
  * [func (this *ListCardItem) SetImage(image string) *ListCardItem](#ListCardItem.SetImage)
  * [func (this *ListCardItem) SetTitle(title string) *ListCardItem](#ListCardItem.SetTitle)
  * [func (this *ListCardItem) SetUrl(url string) *ListCardItem](#ListCardItem.SetUrl)
* [type StandardCard](#StandardCard)
  * [func NewStandardCard() *StandardCard](#NewStandardCard)
  * [func (this *StandardCard) SetContent(content string) *StandardCard](#StandardCard.SetContent)
  * [func (this *StandardCard) SetImage(image string) *StandardCard](#StandardCard.SetImage)
  * [func (this *StandardCard) SetTitle(title string) *StandardCard](#StandardCard.SetTitle)
* [type TextCard](#TextCard)
  * [func NewTextCard(content string) *TextCard](#NewTextCard)
  * [func (this *TextCard) SetContent(content string)](#TextCard.SetContent)


#### <a name="pkg-files">Package files</a>
[base.go](/src/github.com/dueros/bot-sdk-go/bot/card/base.go) [image.go](/src/github.com/dueros/bot-sdk-go/bot/card/image.go) [link_account.go](/src/github.com/dueros/bot-sdk-go/bot/card/link_account.go) [list.go](/src/github.com/dueros/bot-sdk-go/bot/card/list.go) [standard.go](/src/github.com/dueros/bot-sdk-go/bot/card/standard.go) [text.go](/src/github.com/dueros/bot-sdk-go/bot/card/text.go) 






## <a name="BaseCard">type</a> [BaseCard](/src/target/base.go?s=14:65#L3)
``` go
type BaseCard struct {
    Type string `json:"type"`
}
```









## <a name="ImageCard">type</a> [ImageCard](/src/target/image.go?s=14:89#L3)
``` go
type ImageCard struct {
    BaseCard
    List []map[string]string `json:"list"`
}
```






### <a name="NewImageCard">func</a> [NewImageCard](/src/target/image.go?s=91:121#L8)
``` go
func NewImageCard() *ImageCard
```




### <a name="ImageCard.AddItem">func</a> (\*ImageCard) [AddItem](/src/target/image.go?s=183:254#L14)
``` go
func (this *ImageCard) AddItem(src string, thumbnail string) *ImageCard
```



## <a name="LinkAccountCard">type</a> [LinkAccountCard](/src/target/link_account.go?s=14:55#L3)
``` go
type LinkAccountCard struct {
    BaseCard
}
```






### <a name="NewLinkAccountCard">func</a> [NewLinkAccountCard](/src/target/link_account.go?s=57:99#L7)
``` go
func NewLinkAccountCard() *LinkAccountCard
```




## <a name="ListCard">type</a> [ListCard](/src/target/list.go?s=654:724#L35)
``` go
type ListCard struct {
    List []*ListCardItem `json:"list"`
    BaseCard
}
```






### <a name="NewListCard">func</a> [NewListCard](/src/target/list.go?s=726:754#L40)
``` go
func NewListCard() *ListCard
```




### <a name="ListCard.AddItem">func</a> (\*ListCard) [AddItem](/src/target/list.go?s=814:873#L46)
``` go
func (this *ListCard) AddItem(item *ListCardItem) *ListCard
```



## <a name="ListCardItem">type</a> [ListCardItem](/src/target/list.go?s=14:168#L3)
``` go
type ListCardItem struct {
    Title   string `json:"title"`
    Content string `json:"content"`
    Image   string `json:"content"`
    Url     string `json:"url"`
}
```






### <a name="NewListCardItem">func</a> [NewListCardItem](/src/target/list.go?s=170:206#L10)
``` go
func NewListCardItem() *ListCardItem
```




### <a name="ListCardItem.SetContent">func</a> (\*ListCardItem) [SetContent](/src/target/list.go?s=351:417#L20)
``` go
func (this *ListCardItem) SetContent(content string) *ListCardItem
```



### <a name="ListCardItem.SetImage">func</a> (\*ListCardItem) [SetImage](/src/target/list.go?s=460:522#L25)
``` go
func (this *ListCardItem) SetImage(image string) *ListCardItem
```



### <a name="ListCardItem.SetTitle">func</a> (\*ListCardItem) [SetTitle](/src/target/list.go?s=250:312#L15)
``` go
func (this *ListCardItem) SetTitle(title string) *ListCardItem
```



### <a name="ListCardItem.SetUrl">func</a> (\*ListCardItem) [SetUrl](/src/target/list.go?s=561:619#L30)
``` go
func (this *ListCardItem) SetUrl(url string) *ListCardItem
```



## <a name="StandardCard">type</a> [StandardCard](/src/target/standard.go?s=14:147#L3)
``` go
type StandardCard struct {
    Title   string `json:"title"`
    Content string `json:"content"`
    Image   string `json:"image"`
    BaseCard
}
```






### <a name="NewStandardCard">func</a> [NewStandardCard](/src/target/standard.go?s=149:185#L10)
``` go
func NewStandardCard() *StandardCard
```




### <a name="StandardCard.SetContent">func</a> (\*StandardCard) [SetContent](/src/target/standard.go?s=354:420#L21)
``` go
func (this *StandardCard) SetContent(content string) *StandardCard
```



### <a name="StandardCard.SetImage">func</a> (\*StandardCard) [SetImage](/src/target/standard.go?s=463:525#L26)
``` go
func (this *StandardCard) SetImage(image string) *StandardCard
```



### <a name="StandardCard.SetTitle">func</a> (\*StandardCard) [SetTitle](/src/target/standard.go?s=253:315#L16)
``` go
func (this *StandardCard) SetTitle(title string) *StandardCard
```



## <a name="TextCard">type</a> [TextCard](/src/target/text.go?s=14:81#L3)
``` go
type TextCard struct {
    BaseCard
    Content string `json:"content"`
}
```






### <a name="NewTextCard">func</a> [NewTextCard](/src/target/text.go?s=83:125#L8)
``` go
func NewTextCard(content string) *TextCard
```




### <a name="TextCard.SetContent">func</a> (\*TextCard) [SetContent](/src/target/text.go?s=210:258#L15)
``` go
func (this *TextCard) SetContent(content string)
```







- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
