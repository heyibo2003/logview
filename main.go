package main

import (
	_ "LogView/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)
func main() {
	log := logs.NewLogger()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
	log.SetLogger(logs.AdapterConsole,`{"level":7}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
	logs.SetLogger(logs.AdapterFile,`{"filename":"logview.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	beego.Run()
}

