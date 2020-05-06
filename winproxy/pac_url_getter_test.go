package winproxy

import (
	"testing"
)

func TestGetWinPACUrl(t *testing.T) {
	url, err := GetWinPACUrl()
	if err != nil {
		t.Errorf("get pac url error by %s", err.Error())
		t.FailNow()
	}

	// return like http://127.0.0.1:41082/pac?XXXXXXXXXX=
	t.Logf("get win url %s", url)
}
