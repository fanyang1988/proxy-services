package winproxy

import "testing"

func TestGetPAC(t *testing.T) {
	data, err := GetPAC()

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("pac:")
	t.Logf("%s", string(data))
}
