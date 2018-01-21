package nlu

import (
	"../data"
)

type Intent struct {
	data               data.Intent
	Name               string
	ConfirmationStatus string
}

func NewIntent(intent data.Intent) *Intent {
	return &Intent{
		data: intent,
		Name: intent.Name,

		ConfirmationStatus: intent.ConfirmationStatus,
	}
}

func (this *Intent) GetSlot(name string) (*data.Slot, bool) {
	slot, ok := this.data.Slots[name]
	if ok {
		return &slot, true
	}
	return nil, false
}

func (this *Intent) GetSlotValue(name string) (string, bool) {
	slot, ok := this.GetSlot(name)
	if ok {
		return slot.Value, true
	}
	return "", false
}

func (this *Intent) GetSlotStatus(name string) (string, bool) {
	slot, ok := this.GetSlot(name)
	if ok {
		return slot.ConfirmationStatus, true
	}
	return "", false

}

func (this *Intent) SetSlotValue(name string, value string) bool {
	slot, ok := this.GetSlot(name)
	if ok {
		slot.Value = value
		return true
	}
	return false
}
