package data

type SessionResponse struct {
	Attributes map[string]string `json:"attributes"`
}

type ContextResponse struct {
	Intent Intent `json:"intent"`
}

type DialogDirective struct {
	Type          string `json:"type"`
	SlotToElicit  string `json:"slotToElicit,omitempty"`
	SlotToConfirm string `json:"slotToConfirm,omitempty"`
	UpdatedIntent Intent `json:"updatedIntent"`
}

type Speech struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	Ssml string `json:"ssml,omitempty"`
}
