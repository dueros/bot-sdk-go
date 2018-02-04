package data

type DialogDirective struct {
	Type          string `json:"type"`
	SlotToElicit  string `json:"slotToElicit,omitempty"`
	SlotToConfirm string `json:"slotToConfirm,omitempty"`
	UpdatedIntent Intent `json:"updatedIntent"`
}

type Speech struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Ssml string `json:"ssml"`
}
