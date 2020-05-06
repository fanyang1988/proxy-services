package winproxy

import (
	"errors"
	"net"
	"strings"

	"github.com/cihub/seelog"
)

// isSkipInterfaces this not current ip for win
func isSkipInterfaces(name string) bool {
	datas := []string{
		"VMware",
		"vEthernet", // for hyper-v
	}

	for _, d := range datas {
		if strings.Contains(name, d) {
			return true
		}
	}

	return false
}

// GetIP get current IP address
func GetIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, inter := range interfaces {
		addrs, err := inter.Addrs()
		if err != nil {
			continue
		}

		if isSkipInterfaces(inter.Name) {
			seelog.Infof("skip no interfaces %s", inter.Name)
			continue
		}

		seelog.Debugf("net address %s %v", inter.Name, addrs)

		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String(), nil
				}
			}
		}
	}

	return "", errors.New("no ip found")
}
