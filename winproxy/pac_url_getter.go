package winproxy

import (
	"github.com/cihub/seelog"
	"github.com/fanyang1988/go-get-proxied/winhttp"
)

const (
	userAgent = "ir_agent"
)

var (
	resolveTimeout int = 5000
	connectTimeout int = 5000
	sendTimeout    int = 20000
	receiveTimeout int = 20000
)

// GetWinPACUrl get pac url cfg in win10
func GetWinPACUrl() (string, error) {
	h, err := winhttp.Open(
		winhttp.StringToLpwstr(userAgent),
		winhttp.WINHTTP_ACCESS_TYPE_NO_PROXY,
		winhttp.StringToLpwstr(""),
		winhttp.StringToLpwstr(""),
		0)

	if err != nil {
		return "", err
	}

	defer func() {
		if err := winhttp.CloseHandle(h); err != nil {
			seelog.Errorf("[proxy.Provider.closeHandle] Failed to close handle \"%d\": %s\n", h, err)
		}
	}()

	err = winhttp.SetTimeouts(h, resolveTimeout, connectTimeout, sendTimeout, receiveTimeout)
	if err != nil {
		return "", err
	}

	res, err := winhttp.GetIEProxyConfigForCurrentUser()
	if err != nil {
		return "", err
	}

	return winhttp.LpwstrToString(res.LpszAutoConfigUrl), nil
}
