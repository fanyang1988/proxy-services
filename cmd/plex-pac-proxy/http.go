package main

import (
	"fmt"
	"net/http"

	"github.com/cihub/seelog"
	"github.com/fanyang1988/proxy-services/plexpac"
)

var pacData *plexpac.PacDataGetter

func pac(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pac\n")
}

func startPacServer(url, ip string) error {
	seelog.Infof("start pac server by %s with ip %s", url, ip)

	pacData = plexpac.NewPacDataGetter(ip)

	http.HandleFunc("/pac", pac)
	http.ListenAndServe(url, nil)
	return nil
}
