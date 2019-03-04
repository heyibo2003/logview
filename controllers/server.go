package controllers

import (
	"strings"
	"bufio"
	"io"
	"LogView/models"
	"fmt"
	"strconv"
)

type execlist struct {
	Id        int
	Domain    string
	Script    string
	Checkscript string
	Path      string
	Port      int
	Softname  string
	Softpath  string
	Machine   string
	User      string
	Passwd    string
	Hostport  int
}

type ServerController struct {
	BaseController
}

func (c *ServerController) Prepare() {
	c.BaseController.Prepare()
	c.checkLogin()
}

func (c *ServerController) Get() {
	var Software models.Software
	dataList,err :=Software.GetPerjectList()
	if err == nil {
		c.Data["PList"] = dataList
	}
	c.TplName = "list.html"
}

func (c *ServerController) GetEnv(){
	PerjectId,_ :=c.GetInt("PerjectId")
	var Software models.Software
	datalist,err :=Software.GetEnvList(PerjectId)
	if err == nil {
		//c.Data["DList"] = domains
		c.Data["json"] = datalist
	}
	c.ServeJSON()
}
func (c *ServerController) GetApplications(){
	PerjectId,_ :=c.GetInt("PerjectId")
	EnvId,_ :=c.GetInt("EnvId")
	var Software models.Software
	datalist,err :=Software.GetSoftwareByPerjectIdAndEnvId(PerjectId,EnvId)
	if err == nil {
		//c.Data["DList"] = domains
		c.Data["json"] = datalist
	}
	c.ServeJSON()
}
func (c *ServerController) GetHosts()  {
	PerjectId,_ :=c.GetInt("PerjectId")
	EnvId,_ :=c.GetInt("EnvId")
	Domain :=c.GetString("domain")
	var Environmentals models.Environmentals
	datalist,err :=Environmentals.GetMachineBydDomainAndEnvIdAndPerjectId(PerjectId,EnvId,Domain)
	if err == nil {
		//c.Data["DList"] = domains
		c.Data["json"] = datalist
	}
	c.ServeJSON()
}


func (c * ServerController) GetExecLogs() {
	PerjectId,_ :=c.GetInt("PerjectId")
	EnvId,_ :=c.GetInt("EnvId")
	Domain :=c.GetString("Domain")
	Machines :=c.GetString("Machine")
	ExecName :=c.GetString("ExecName")
	var Software models.Software
	var  Remote models.Remote
	var  pro execlist
	var  CMD string = ""
	dataList, err := Software.GetSoftwareByEnvironmentals(PerjectId,EnvId,Domain,Machines)
	if err == nil {
		c.Data["List"] = dataList
		for _, v := range dataList {
			pro.Id =v.Id
			pro.Domain =v.Domain
			pro.Script = v.Script
			pro.Checkscript = v.Checkscript
			pro.Path  = v.Path
			pro.Port  = v.Port
			pro.Softname = v.Softname
			pro.Softpath = v.Softpath
			pro.Machine = v.Machine
			pro.User  =v.User
			pro.Passwd =v.Passwd
			pro.Hostport  =v.Hostport
		}
	}
    port :=strconv.Itoa(pro.Port)
	if ExecName == "restart" {
		CheckName :="sh "+pro.Path+""+pro.Checkscript+" "+pro.Softname+" "+port+";cd "+pro.Softpath+";cd "+pro.Softpath+";tail -n 100 stdout.log"
		CMD =fmt.Sprintf("%s;%s;%s", "cd "+pro.Softpath+";sh "+pro.Path+""+pro.Script+" "+pro.Softname+" "+port+" "+ExecName+"","sleep 30s",CheckName)
	}else if ExecName == "stop"{
		CheckName :="sh "+pro.Path+""+pro.Checkscript+" "+pro.Softname+" "+port+""
		CMD =fmt.Sprintf("%s;%s", "cd "+pro.Softpath+";sh "+pro.Path+""+pro.Script+" "+pro.Softname+" "+port+" "+ExecName+"",CheckName)
	}else if ExecName == "start" {
		CheckName :="sh "+pro.Path+""+pro.Checkscript+" "+pro.Softname+" "+port+";cd "+pro.Softpath+";cd "+pro.Softpath+";tail -n 100 stdout.log"
		CMD =fmt.Sprintf("%s;%s;%s", "cd "+pro.Softpath+";sh "+pro.Path+""+pro.Script+" "+pro.Softname+" "+port+" "+ExecName+"","sleep 30s",CheckName)
	}else {
		CMD ="sh "+pro.Path+""+pro.Checkscript+" "+pro.Softname+" "+port+""
		//fmt.Println("cmd:",CMD)
	}
	stdout,err :=Remote.GetSshClientByCMD(pro.User,pro.Passwd,pro.Machine,pro.Hostport,CMD)
	data :=strings.NewReader(stdout)
	reader := bufio.NewReader(data)
	var index int
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		if 0 == len(line) || line == "\r\n" {
			continue
		}
		index++
		line = strings.Replace(line, "\n", "", -1)
		msg := models.Message{Message:line}
		broadcast <- msg
	}
	c.Data["json"] = "成功！"
	c.ServeJSON()
}