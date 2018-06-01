package audio_player

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

func NewPlayerInfo() *PlayerInfo {
	playerInfo := &PlayerInfo{}

	playerInfo.Content.AudioItemType = AUDIO_TYPE_MUSIC

	return playerInfo
}

func (this *PlayerInfo) SetTitle(title string) *PlayerInfo {
	this.Content.Title = title
	return this
}

func (this *PlayerInfo) SetTitleSubtext1(text string) *PlayerInfo {
	this.Content.TitleSubtext1 = text
	return this
}

func (this *PlayerInfo) SetTitleSubtext2(text string) *PlayerInfo {
	this.Content.TitleSubtext2 = text
	return this
}

func (this *PlayerInfo) SetLyric(url string) *PlayerInfo {
	this.Content.Lyric.Url = url
	this.Content.Lyric.Format = LYRIC_FORMAT_LRC
	return this
}

func (this *PlayerInfo) SetMediaLengthInMs(length uint64) *PlayerInfo {
	this.Content.MediaLengthInMilliseconds = length
	return this
}

func (this *PlayerInfo) SetArt(src string) *PlayerInfo {
	this.Content.Art.Src = src
	return this
}

func (this *PlayerInfo) SetProviderName(name string) *PlayerInfo {
	this.Content.Provider.Name = name
	return this
}

func (this *PlayerInfo) SetProviderLogo(logo string) *PlayerInfo {
	this.Content.Provider.Logo.Src = logo
	return this
}
