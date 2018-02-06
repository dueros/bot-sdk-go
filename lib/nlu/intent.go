package nlu

import (
	"github.com/dueros/bot-sdk-go/lib/data"
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

func (this *Intent) GetSlot(name string) *data.Slot {
	slot, ok := this.data.Slots[name]
	if ok {
		return &slot
	}
	return nil
}

func (this *Intent) GetSlotValue(name string) string {
	slot := this.GetSlot(name)
	if slot != nil {
		return slot.Value
	}
	return ""
}

func (this *Intent) GetSlotStatus(name string) string {
	slot := this.GetSlot(name)
	if slot != nil {
		return slot.ConfirmationStatus
	}
	return ""

}

func (this *Intent) SetSlotValue(name string, value string) bool {
	slot := this.GetSlot(name)
	if slot != nil {
		slot.Value = value
		return true
	} else {
		// TODO new a slot
	}
	return false
}

func (this *Intent) GetData() data.Intent {
	return this.data
}
