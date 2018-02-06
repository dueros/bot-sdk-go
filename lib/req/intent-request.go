package req

func (this *IntentRequest) GetIntentName() (string, bool) {
	return this.Dialog.GetIntentName()
}

func (this *IntentRequest) IsDialogStateCompleted() bool {
	return true
}

func (this *IntentRequest) GetQuery() string {
	query, _ := this.Dialog.GetQuery()
	return query
}
