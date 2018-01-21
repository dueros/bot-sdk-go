package req

func (this *IntentRequest) GetIntentName() (string, bool) {
	return this.Nlu.GetIntentName()
}

func (this *IntentRequest) IsDialogStateCompleted() bool {
	return true
}
