package plexpac

import (
	"fmt"
	"strings"
)

func changePAC(ip string, data []byte) []byte {
	return []byte(
		strings.Replace(
			string(data),
			"PROXY 127.0.0.1",
			fmt.Sprintf("PROXY %s", ip),
			1),
	)
}
