

# template
`import "github.com/dueros/bot-sdk-go/bot/directive/display/template"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type BaseTemplate](#BaseTemplate)
  * [func (this *BaseTemplate) SetBackgroundImage(background *Image)](#BaseTemplate.SetBackgroundImage)
  * [func (this *BaseTemplate) SetBackgroundImageUrl(url string)](#BaseTemplate.SetBackgroundImageUrl)
  * [func (this *BaseTemplate) SetTitle(title string)](#BaseTemplate.SetTitle)
* [type BodyTemplate1](#BodyTemplate1)
  * [func NewBodyTemplate1() *BodyTemplate1](#NewBodyTemplate1)
  * [func (this *BodyTemplate1) SetContentPosition(position string) *BodyTemplate1](#BodyTemplate1.SetContentPosition)
  * [func (this *BodyTemplate1) SetPlainContent(content string) *BodyTemplate1](#BodyTemplate1.SetPlainContent)
* [type BodyTemplate2](#BodyTemplate2)
  * [func NewBodyTemplate2() *BodyTemplate2](#NewBodyTemplate2)
* [type BodyTemplate3](#BodyTemplate3)
  * [func NewBodyTemplate3() *BodyTemplate3](#NewBodyTemplate3)
* [type BodyTemplate4](#BodyTemplate4)
  * [func NewBodyTemplate4() *BodyTemplate4](#NewBodyTemplate4)
* [type BodyTemplate5](#BodyTemplate5)
  * [func NewBodyTemplate5() *BodyTemplate5](#NewBodyTemplate5)
  * [func (this *BodyTemplate5) AddImage(image *Image) *BodyTemplate5](#BodyTemplate5.AddImage)
* [type Image](#Image)
  * [func NewImage(url string) *Image](#NewImage)
  * [func (this *Image) SetHeight(height int) *Image](#Image.SetHeight)
  * [func (this *Image) SetWidth(width int) *Image](#Image.SetWidth)
* [type ListItem](#ListItem)
  * [func NewListItem() *ListItem](#NewListItem)
  * [func (this *ListItem) SetImage(image *Image) *ListItem](#ListItem.SetImage)
  * [func (this *ListItem) SetImageUrl(url string) *ListItem](#ListItem.SetImageUrl)
  * [func (this *ListItem) SetPlainPrimaryText(text string) *ListItem](#ListItem.SetPlainPrimaryText)
  * [func (this *ListItem) SetPlainSecondaryText(text string) *ListItem](#ListItem.SetPlainSecondaryText)
  * [func (this *ListItem) SetPlainTertiary(text string) *ListItem](#ListItem.SetPlainTertiary)
* [type ListTemplate](#ListTemplate)
  * [func (this *ListTemplate) AddItem(listItem *ListItem)](#ListTemplate.AddItem)
* [type ListTemplate1](#ListTemplate1)
  * [func NewListTemplate1() *ListTemplate1](#NewListTemplate1)
* [type ListTemplate2](#ListTemplate2)
  * [func NewListTemplate2() *ListTemplate2](#NewListTemplate2)
* [type Text](#Text)
  * [func NewText(textType, text string) *Text](#NewText)
* [type TextImageTemplate](#TextImageTemplate)
  * [func (this *TextImageTemplate) SetImage(image *Image)](#TextImageTemplate.SetImage)
  * [func (this *TextImageTemplate) SetImageUrl(url string)](#TextImageTemplate.SetImageUrl)
  * [func (this *TextImageTemplate) SetPlainContent(content string)](#TextImageTemplate.SetPlainContent)


#### <a name="pkg-files">Package files</a>
[base.go](/src/github.com/dueros/bot-sdk-go/bot/directive/display/template/base.go) [body_template1.go](/src/github.com/dueros/bot-sdk-go/bot/directive/display/template/body_template1.go) [body_template2_3_4.go](/src/github.com/dueros/bot-sdk-go/bot/directive/display/template/body_template2_3_4.go) [body_template5.go](/src/github.com/dueros/bot-sdk-go/bot/directive/display/template/body_template5.go) [list.go](/src/github.com/dueros/bot-sdk-go/bot/directive/display/template/list.go) [list_template1_2.go](/src/github.com/dueros/bot-sdk-go/bot/directive/display/template/list_template1_2.go) [text_image.go](/src/github.com/dueros/bot-sdk-go/bot/directive/display/template/text_image.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    PLAIN_TEXT = "PlainText"
    RICH_TEXT  = "RichText"
)
```
``` go
const (
    BOTTOM_LEFT_POSITION = "BOTTOM-LEFT"
    CENTER               = "CENTER"
    TOP_LEFT             = "TOP-LEFT"
)
```




## <a name="BaseTemplate">type</a> [BaseTemplate](/src/target/base.go?s=214:404#L17)
``` go
type BaseTemplate struct {
    directive.BaseDirective
    Token           string `json:"token"`
    BackgroundImage *Image `json:"backgroundImage,omitempty"`
    Title           string `json:"title"`
}
```









### <a name="BaseTemplate.SetBackgroundImage">func</a> (\*BaseTemplate) [SetBackgroundImage](/src/target/base.go?s=585:648#L32)
``` go
func (this *BaseTemplate) SetBackgroundImage(background *Image)
```



### <a name="BaseTemplate.SetBackgroundImageUrl">func</a> (\*BaseTemplate) [SetBackgroundImageUrl](/src/target/base.go?s=480:539#L28)
``` go
func (this *BaseTemplate) SetBackgroundImageUrl(url string)
```



### <a name="BaseTemplate.SetTitle">func</a> (\*BaseTemplate) [SetTitle](/src/target/base.go?s=406:454#L24)
``` go
func (this *BaseTemplate) SetTitle(title string)
```



## <a name="BodyTemplate1">type</a> [BodyTemplate1](/src/target/body_template1.go?s=264:417#L15)
``` go
type BodyTemplate1 struct {
    BaseTemplate
    TextContent struct {
        Position string `json:"position"`
        Text     *Text  `json:"text"`
    } `json:"content"`
}
```






### <a name="NewBodyTemplate1">func</a> [NewBodyTemplate1](/src/target/body_template1.go?s=419:457#L23)
``` go
func NewBodyTemplate1() *BodyTemplate1
```




### <a name="BodyTemplate1.SetContentPosition">func</a> (\*BodyTemplate1) [SetContentPosition](/src/target/body_template1.go?s=659:736#L31)
``` go
func (this *BodyTemplate1) SetContentPosition(position string) *BodyTemplate1
```



### <a name="BodyTemplate1.SetPlainContent">func</a> (\*BodyTemplate1) [SetPlainContent](/src/target/body_template1.go?s=874:947#L41)
``` go
func (this *BodyTemplate1) SetPlainContent(content string) *BodyTemplate1
```



## <a name="BodyTemplate2">type</a> [BodyTemplate2](/src/target/body_template2_3_4.go?s=18:66#L3)
``` go
type BodyTemplate2 struct {
    TextImageTemplate
}
```






### <a name="NewBodyTemplate2">func</a> [NewBodyTemplate2](/src/target/body_template2_3_4.go?s=68:106#L7)
``` go
func NewBodyTemplate2() *BodyTemplate2
```




## <a name="BodyTemplate3">type</a> [BodyTemplate3](/src/target/body_template2_3_4.go?s=250:298#L14)
``` go
type BodyTemplate3 struct {
    TextImageTemplate
}
```






### <a name="NewBodyTemplate3">func</a> [NewBodyTemplate3](/src/target/body_template2_3_4.go?s=300:338#L18)
``` go
func NewBodyTemplate3() *BodyTemplate3
```




## <a name="BodyTemplate4">type</a> [BodyTemplate4](/src/target/body_template2_3_4.go?s=482:530#L25)
``` go
type BodyTemplate4 struct {
    TextImageTemplate
}
```






### <a name="NewBodyTemplate4">func</a> [NewBodyTemplate4](/src/target/body_template2_3_4.go?s=532:570#L29)
``` go
func NewBodyTemplate4() *BodyTemplate4
```




## <a name="BodyTemplate5">type</a> [BodyTemplate5](/src/target/body_template5.go?s=18:94#L3)
``` go
type BodyTemplate5 struct {
    BaseTemplate
    Images []*Image `json:"images"`
}
```






### <a name="NewBodyTemplate5">func</a> [NewBodyTemplate5](/src/target/body_template5.go?s=96:134#L8)
``` go
func NewBodyTemplate5() *BodyTemplate5
```




### <a name="BodyTemplate5.AddImage">func</a> (\*BodyTemplate5) [AddImage](/src/target/body_template5.go?s=236:300#L14)
``` go
func (this *BodyTemplate5) AddImage(image *Image) *BodyTemplate5
```



## <a name="Image">type</a> [Image](/src/target/base.go?s=689:849#L36)
``` go
type Image struct {
    Url          string `json:"url"`
    WidthPixels  int    `json:"widthPixels,omitempty"`
    HeightPixels int    `json:"heightPixels,omitempty"`
}
```






### <a name="NewImage">func</a> [NewImage](/src/target/base.go?s=851:883#L42)
``` go
func NewImage(url string) *Image
```




### <a name="Image.SetHeight">func</a> (\*Image) [SetHeight](/src/target/base.go?s=1029:1076#L53)
``` go
func (this *Image) SetHeight(height int) *Image
```



### <a name="Image.SetWidth">func</a> (\*Image) [SetWidth](/src/target/base.go?s=939:984#L48)
``` go
func (this *Image) SetWidth(width int) *Image
```



## <a name="ListItem">type</a> [ListItem](/src/target/list.go?s=18:310#L3)
``` go
type ListItem struct {
    Token       string `json:"token"`
    Image       *Image `json:"image,omitempty"`
    TextContent struct {
        PrimaryText   *Text `json:"primaryText"`
        SecondaryText *Text `json:"secondaryText"`
        TertiaryText  *Text `json:"tertiaryText,omitempty"`
    } `json:"textContent"`
}
```






### <a name="NewListItem">func</a> [NewListItem](/src/target/list.go?s=312:340#L13)
``` go
func NewListItem() *ListItem
```




### <a name="ListItem.SetImage">func</a> (\*ListItem) [SetImage](/src/target/list.go?s=380:434#L18)
``` go
func (this *ListItem) SetImage(image *Image) *ListItem
```



### <a name="ListItem.SetImageUrl">func</a> (\*ListItem) [SetImageUrl](/src/target/list.go?s=473:528#L23)
``` go
func (this *ListItem) SetImageUrl(url string) *ListItem
```



### <a name="ListItem.SetPlainPrimaryText">func</a> (\*ListItem) [SetPlainPrimaryText](/src/target/list.go?s=587:651#L28)
``` go
func (this *ListItem) SetPlainPrimaryText(text string) *ListItem
```



### <a name="ListItem.SetPlainSecondaryText">func</a> (\*ListItem) [SetPlainSecondaryText](/src/target/list.go?s=728:794#L33)
``` go
func (this *ListItem) SetPlainSecondaryText(text string) *ListItem
```



### <a name="ListItem.SetPlainTertiary">func</a> (\*ListItem) [SetPlainTertiary](/src/target/list.go?s=873:934#L38)
``` go
func (this *ListItem) SetPlainTertiary(text string) *ListItem
```



## <a name="ListTemplate">type</a> [ListTemplate](/src/target/list.go?s=1012:1096#L43)
``` go
type ListTemplate struct {
    BaseTemplate
    ListItems []*ListItem `json:"listItems"`
}
```









### <a name="ListTemplate.AddItem">func</a> (\*ListTemplate) [AddItem](/src/target/list.go?s=1098:1151#L48)
``` go
func (this *ListTemplate) AddItem(listItem *ListItem)
```



## <a name="ListTemplate1">type</a> [ListTemplate1](/src/target/list_template1_2.go?s=18:61#L3)
``` go
type ListTemplate1 struct {
    ListTemplate
}
```






### <a name="NewListTemplate1">func</a> [NewListTemplate1](/src/target/list_template1_2.go?s=63:101#L7)
``` go
func NewListTemplate1() *ListTemplate1
```




## <a name="ListTemplate2">type</a> [ListTemplate2](/src/target/list_template1_2.go?s=245:288#L14)
``` go
type ListTemplate2 struct {
    ListTemplate
}
```






### <a name="NewListTemplate2">func</a> [NewListTemplate2](/src/target/list_template1_2.go?s=290:328#L18)
``` go
func NewListTemplate2() *ListTemplate2
```




## <a name="Text">type</a> [Text](/src/target/base.go?s=1123:1197#L58)
``` go
type Text struct {
    Type string `json:"type"`
    Text string `json:"text"`
}
```






### <a name="NewText">func</a> [NewText](/src/target/base.go?s=1199:1240#L63)
``` go
func NewText(textType, text string) *Text
```




## <a name="TextImageTemplate">type</a> [TextImageTemplate](/src/target/text_image.go?s=18:129#L3)
``` go
type TextImageTemplate struct {
    BaseTemplate
    Content *Text  `json:"content"`
    Image   *Image `json:"image"`
}
```









### <a name="TextImageTemplate.SetImage">func</a> (\*TextImageTemplate) [SetImage](/src/target/text_image.go?s=364:417#L19)
``` go
func (this *TextImageTemplate) SetImage(image *Image)
```



### <a name="TextImageTemplate.SetImageUrl">func</a> (\*TextImageTemplate) [SetImageUrl](/src/target/text_image.go?s=258:312#L14)
``` go
func (this *TextImageTemplate) SetImageUrl(url string)
```



### <a name="TextImageTemplate.SetPlainContent">func</a> (\*TextImageTemplate) [SetPlainContent](/src/target/text_image.go?s=131:193#L9)
``` go
func (this *TextImageTemplate) SetPlainContent(content string)
```







- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
