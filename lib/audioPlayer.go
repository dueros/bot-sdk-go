package lib

type AudioPlayer struct {
	Bot
}

func NewAudioPlayer(rawRequest string) *AudioPlayer {
	return &AudioPlayer{Bot{
		intentHandler: make(map[string]func() ResponseData),
		eventHandler:  make(map[string]func() ResponseData),
		Request:       NewRequest(rawRequest),
	}}
}

func (this *AudioPlayer) OnPlaybackStarted(fn func() ResponseData) {
	this.AddEventListener(AUDIO_PLAYER_PLAYBACK_STARTED, fn)
}
