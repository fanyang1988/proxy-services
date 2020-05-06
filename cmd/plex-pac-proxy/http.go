package main

import (
	"net/http"

	"github.com/cihub/seelog"
	"github.com/fanyang1988/proxy-services/plexpac"
)

var pacData *plexpac.PacDataGetter

func pac(w http.ResponseWriter, req *http.Request) {
	//Content-Disposition: attachment;filename=pac
	//Access-Control-Allow-Origin: *
	//Access-Control-Allow-Headers: Content-Type,X-Requested-With
	//MIME type: application/x-ns-proxy-autoconfig

	w.Header().Set("Content-Disposition", "attachment;filename=pac")
	w.Header().Set("Access-Control-Allow-Origin", "*")                               //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", " Content-Type,X-Requested-With") //header的类型
	w.Header().Set("content-type", "application/x-ns-proxy-autoconfig")              //返回数据格式是json

	w.Write(pacData.GetPAC())
}

func startPacServer(url, ip string) error {
	seelog.Infof("start pac server by %s with ip %s", url, ip)

	pacData = plexpac.NewPacDataGetter(ip)

	http.HandleFunc("/pac", pac)
	http.ListenAndServe(url, nil)
	return nil
}
