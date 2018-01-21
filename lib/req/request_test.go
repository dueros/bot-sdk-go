package req

import "testing"

func TestNewRequest(t *testing.T) {
	request := NewRequest()

	if request.Data.Version != "v2.0" {
		t.Errorf("New Request Error")
	}
}
