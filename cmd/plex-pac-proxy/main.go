package main

import (
	"flag"

	"github.com/cihub/seelog"
	"github.com/fanyang1988/proxy-services/config"
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
	if pacURL != "" {
		pacURL = *url
	}

	startPacServer(pacURL)
}

func startPacServer(url string) error {
	seelog.Infof("start pac server by %s", url)

	return nil
}
