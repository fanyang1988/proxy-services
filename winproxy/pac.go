package winproxy

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// GetPAC get pac datas
func GetPAC() ([]byte, error) {
	url, err := GetWinPACUrl()
	if err != nil {
		return []byte{}, errors.Wrap(err, "get pac url error")
	}

	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "http get %s error", url)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.Wrap(err, "read all error")
	}

	return body, nil
}
