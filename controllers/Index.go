package controllers

import (
	"LogView/models"
	"github.com/astaxie/beego/logs"
	"time"
	"fmt"
	"strings"
	"bufio"
	"io"
)

type MainController struct {
	BaseController
}

func (c *MainController) Prepare() {
	 c.BaseController.Prepare()
     c.checkLogin()
}



func (c *MainController) Get() {
	var Env models.Env
	dataList2,err :=Env.GetAll()
	if err == nil {
		c.Data["EList"] = dataList2
	}
    c.Data["LogsTime"]	= time.Now().Format("2006-01-02")
	c.TplName = "index.html"
}

func (c *MainController) GetDomains() {
	EnvId,_ :=c.GetInt("EnvId")
	var Environmentals models.Environmentals
	domains,err :=Environmentals.GetDomainsByEnv(EnvId)
	if err == nil {
		//c.Data["DList"] = domains
		c.Data["json"] = domains
	}
	logs.Info("dataList-damain :", domains)
	c.ServeJSON()
}

func (c *MainController) GetMachines() {
	EnvId,_ :=c.GetInt("EnvId")
    Domain :=c.GetString("domain")
	var Environmentals models.Environmentals
	data,err :=Environmentals.GetMachineBydDomain(EnvId,Domain)
	if err == nil {
		//c.Data["DList"] = domains
		c.Data["json"] = data
	}
	logs.Info("dataList-machines :", data)
	c.ServeJSON()
}

func (c * MainController) GetLogsName(){
	EnvId,_ :=c.GetInt("EnvId")
	Domain :=c.GetString("domain")
	Machines :=c.GetString("machine")
	var Environmentals models.Environmentals
	data,err :=Environmentals.GetNameByDomainAndMachine(EnvId,Domain,Machines)
	if err == nil {
		//c.Data["DList"] = domains
		c.Data["json"] = data
	}
	logs.Info("dataList-logsname :", data)
	c.ServeJSON()
}
func (c *MainController) GetLogContents (){
	var Remote models.Remote
	var line string =""
	var min string = ""
	var BeginTime string = ""
	var EndTime string = ""
	var LogsNamePath   = ""
	var LogsContent  string= ""
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	EnvId,err :=c.GetInt("EnvId")
	if err != nil {
		logs.Error("EnvId的参数:",EnvId)
		Msg :="EnvId参数错误！"
		LogsContent = fmt.Sprintf("%s:%s", Msg, err)
		c.Data["json"] =  "{\"Message\":\"" + Msg + "\"}"
	}
	Domain :=c.GetString("domain")
	if Domain == "" {
		logs.Error("domain参数:",Domain)
		Msg :="Domain参数错误！"
		LogsContent = fmt.Sprintf("%s:%s", Msg, Domain)
		c.Data["json"] = "{\"Message\":\"" + Msg + "\"}"
	}
	Machines :=c.GetString("machine")
	if Machines == "" {
		logs.Error("Machines参数:",Machines)
		Msg :="Machines参数错误！"
		LogsContent = fmt.Sprintf("%s:%s", Msg, Machines)
		c.Data["json"] = "{\"Message\":\"" + Msg + "\"}"
	}
	LogsName :=c.GetString("logsname")
	if LogsName == "" {
		logs.Error("LogsName参数:",LogsName)
		Msg :="LogsName参数错误！"
		LogsContent = fmt.Sprintf("%s:%s", Msg, LogsName)
		c.Data["json"] = "{\"Message\":\"" + Msg + "\"}"
	}
	LogsLevel :=c.GetString("loglevel")
	if LogsLevel == "" {
		logs.Error("LogsLevel参数:",LogsLevel)
		Msg :="LogsLevel参数错误！"
		LogsContent = fmt.Sprintf("%s:%s", Msg, LogsLevel)
		c.Data["json"] = "{\"Message\":\"" + Msg + "\"}"
	}
	LogsTime := c.GetString("LogsTime")
	t1,err :=time.Parse("2006-01-02", LogsTime)
	t2,err :=time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if LogsTime == "" {
		logs.Error("LogsTime参数:",LogsTime)
		Msg :="日期不能为空！"
		LogsContent = fmt.Sprintf("%s:%s", Msg, LogsTime)
		LogsContent = "{\"Message\":\"" + Msg + "\"}"
	}
	ViewType ,err :=c.GetInt("viewtype")
	if err != nil {
		logs.Error("ViewType的参数:",ViewType)
		Msg :="ViewType参数错误！"
		LogsContent = fmt.Sprintf("%s:%s", Msg, err)
		c.Data["json"] =  "{\"Message\":\"" + Msg + "\"}"
	}else if ViewType == 1{
		line =c.GetString("line")
	}else if ViewType == 2 {
		min =c.GetString("min")
	}else if ViewType == 3 {
		BeginTime =c.GetString("BeginTime")
		EndTime =c.GetString("EndTime")
	}
	var DataList models.Environmentals
	list,err :=DataList.GetEnvironmentalsByFourParameter(EnvId,Domain,Machines,LogsName)
	var path string =""
	var user string=""
	var passwd string=""
	var port int= 0
	if err == nil {
		for _, v := range list {
			path  = v.Path
			user  =v.User
			passwd =v.Passwd
			port  =v.Hostport
		}
	}
	if t1.Equal(t2){
		LogsNamePath =fmt.Sprintf("%s%s", path, LogsName)
	}else if t1.Before(t2){
		if LogsName =="stdout.log"{
			LogsNamePath =fmt.Sprintf("%s%s", path, LogsName)
		}else {
			LogsNamePath =fmt.Sprintf("%s%s.%s.%s", path, strings.Split(LogsName, ".")[0],LogsTime,strings.Split(LogsName, ".")[1])
		}
	}else {
		LogsNamePath =fmt.Sprintf("%s%s", path, LogsName)
	}
	if ViewType == 1 {
      stdout,err :=Remote.GetSshClientByExec(user,passwd,Machines,port,"tail -n "+ line +"  "+ LogsNamePath +"")
		if err !=nil {
			logs.Error("执行Linux命令错误！:",err)
			Msg := "执行Linux命令错误！"
			LogsContent = fmt.Sprintf("%s:%s%s", Msg, err,"\n")
			c.Data["json"] =  "{\"Message\":\"" + Msg + "\"}"

		}else {
			LogsContent = stdout
		}
	}else if ViewType == 2{
		    line = "1"
			E_stdout,err :=Remote.GetSshClientByExec(user,passwd,Machines,port,"tail -n "+ line +"  "+ LogsNamePath +"")
		    fmt.Println("LogsNamePath:",LogsNamePath)
			if err !=nil {
				logs.Error("执行Linux命令错误！:",err)
				Msg := "执行Linux命令错误！"
				LogsContent = fmt.Sprintf("%s:%s%s", Msg, err,"\n")
				c.Data["json"] =  "{\"Message\":\"" + Msg + "\"}"
			}else {
				//从最后line行中捕捉所有的时间
				t := strings.Split(E_stdout, ".")
				LogsLastTime, _ := time.ParseInLocation(timeLayout, t[0], loc)
                if LogsLastTime.Format("2006-01-02 15:04:00") == "0001-01-01 00:00:00" {
					Msg := "日志文件"+LogsName+"中无时间格式，请配置日志文件格式"
					LogsContent = fmt.Sprintf("%s%s", Msg, "\n")
				}else {
					negativeM, _ := time.ParseDuration("-" + min + "m")
					LogsSartTime := LogsLastTime.Add(negativeM)
					CMD := fmt.Sprintf("sed -n  '/%s/,/%s/p' %s", LogsSartTime.Format("2006-01-02 15:04"), LogsLastTime.Format("2006-01-02 15:04"), LogsNamePath)
					fmt.Println("cmd1:", CMD)
					stdout, err := Remote.GetSshClientByExec(user, passwd, Machines, port, CMD)
					if err != nil{
						Msg := "执行Linux命令错误！"+CMD+""
						LogsContent = fmt.Sprintf("%s%s", Msg,"\n")
					}else {
						LogsContent = stdout
					}
				}
			}
	}else {
		line = "1"
		E_stdout,err :=Remote.GetSshClientByExec(user,passwd,Machines,port,"tail -n "+ line +"  "+ LogsNamePath +"")
		if err !=nil {
			logs.Error("执行Linux命令错误！:",err)
			Msg := "执行Linux命令错误！"
			LogsContent = fmt.Sprintf("%s:%s%s", Msg, err,"\n")
			c.Data["json"] =  "{\"Message\":\"" + Msg + "\"}"
		}else {
			t :=strings.Split(E_stdout, ".")
			LogsLastTime, _ := time.ParseInLocation(timeLayout, t[0], loc)
			if LogsLastTime.Format("2006-01-02 15:04:00") == "0001-01-01 00:00:00" {
				Msg := "日志文件"+LogsName+"中无时间格式，请配置日志文件格式"
				LogsContent = fmt.Sprintf("%s%s", Msg, "\n")
			}else{
				l_BeginTime :=fmt.Sprintf("%s %s",LogsTime,BeginTime)
				l_EndTime :=fmt.Sprintf("%s %s",LogsTime,EndTime)
				LogsBeginTime, _ :=time.ParseInLocation(timeLayout, l_BeginTime, loc)
				LogsEndTime, _ :=time.ParseInLocation(timeLayout, l_EndTime, loc)
				t_Last,_ :=time.Parse("2006-01-02 15:04:05", t[0])
				t_Begin,_ :=time.Parse("2006-01-02 15:04:05", l_BeginTime)
				t_End,_ :=time.Parse("2006-01-02 15:04:05", l_EndTime)
				if  t_Begin.After(t_Last){
					Msg := "开始时间比文件最后时间新,无法按条件查找!"
					LogsContent = fmt.Sprintf("%s%s", Msg,"\n")
					c.Data["json"] =  "{\"Message\":\"" + Msg + "\"}"
				}else if t_Begin.Before(t_Last) && t_End.Before(t_Last){
					CMD := fmt.Sprintf("sed -n  '/%s/,/%s/p' %s", LogsBeginTime.Format("2006-01-02 15:04"),  LogsEndTime.Format("2006-01-02 15:04"), LogsNamePath)
					fmt.Println("cmd3:",CMD)
					stdout,err :=Remote.GetSshClientByExec(user,passwd,Machines,port,CMD)
					if err !=nil {
						logs.Error("执行Linux命令错误！:",err)
						Msg := "执行Linux命令错误！"
						LogsContent = fmt.Sprintf("%s:%s%s", Msg, err,"\n")
						c.Data["json"] =  "{\"Message\":\"" + Msg + "\"}"
					}else {
						LogsContent = stdout
					}
				}else if t_Begin.Before(t_Last) && t_End.After(t_Last) {
					CMD := fmt.Sprintf("sed -n  '/%s/,/%s/p' %s", LogsBeginTime.Format("2006-01-02 15:04"),  LogsLastTime.Format("2006-01-02 15:04"), LogsNamePath)
					fmt.Println("cmd3-1:",CMD)
					stdout,err :=Remote.GetSshClientByExec(user,passwd,Machines,port,CMD)
					if err !=nil {
						logs.Error("执行Linux命令错误！:",err)
						Msg := "执行Linux命令错误！"
						LogsContent = fmt.Sprintf("%s:%s%s", Msg, err,"\n")
						c.Data["json"] =  "{\"Message\":\"" + Msg + "\"}"
					}else {
						LogsContent = stdout
					}
				}else {
					CMD := fmt.Sprintf("sed -n  '/%s/,/%s/p' %s", LogsBeginTime.Format("2006-01-02 15:04"),  LogsLastTime.Format("2006-01-02 15:04"), LogsNamePath)
					fmt.Println("cmd3-2:",CMD)
					stdout,err :=Remote.GetSshClientByExec(user,passwd,Machines,port,CMD)
					if err !=nil {
						logs.Error("执行Linux命令错误！:",err)
						Msg := "执行Linux命令错误！"
						LogsContent = fmt.Sprintf("%s:%s%s", Msg, err,"\n")
						c.Data["json"] =  "{\"Message\":\"" + Msg + "\"}"
					}else {
						LogsContent = stdout
					}
				}
			}
		}
	}
	data :=strings.NewReader(LogsContent)
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
		switch LogsLevel {
		    case "debug":
		    case "info":
				if strings.Contains(line,"debug")||  strings.Contains(line,"DEBUG"){
					continue
				}
		    case "warn":
				if (strings.Contains(line,"info") || strings.Contains(line,"debug"))|| (strings.Contains(line,"INFO") || strings.Contains(line,"DEBUG")){
					continue
				}
		    case "error":
				if (strings.Contains(line,"debug")|| strings.Contains(line,"warn")|| strings.Contains(line,"info"))||(strings.Contains(line,"DEBUG")|| strings.Contains(line,"WARN") || strings.Contains(line,"INFO")){
					continue
				}
				fmt.Println("12",1)
		}
		index++
		line = strings.Replace(line, "\n", "", -1)
		//fmt.Println("123",line)
		msg := models.Message{Message:line}
		broadcast <- msg
	}
	if index == 0 {
		line = "没有相应的日志！请重新搜索"
	}
	//fmt.Println("123",line)
	msg := models.Message{Message:line}
	broadcast <- msg
	c.Data["json"] = "成功！"
	c.ServeJSON()
}