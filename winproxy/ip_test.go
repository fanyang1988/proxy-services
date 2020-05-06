package winproxy

import "testing"

func TestGetIP(t *testing.T) {
	ip, err := GetIP()

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("ip %s", ip)
}
