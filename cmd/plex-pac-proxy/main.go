package main

import (
	"flag"

	"github.com/cihub/seelog"
	"github.com/fanyang1988/proxy-services/config"
	"github.com/fanyang1988/proxy-services/winproxy"
)

var (
	cfgPath = flag.String("cfg", ".\\cfg.json", "config file path")
	url     = flag.String("u", "", "pac url for bind")
)

func main() {
	defer seelog.Flush()
	flag.Parse()

	cfg := config.LoadCfg(*cfgPath)
	pacURL := cfg.URL
	if pacURL == "" {
		pacURL = *url
	}

	ip := cfg.IP
	if ip == "" {
		var err error
		ip, err = winproxy.GetIP()
		if err != nil {
			seelog.Errorf("get ip error, cfg is nil by %s", err.Error())
			return
		}
	}

	startPacServer(pacURL, ip)
}
