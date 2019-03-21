package controllers
import (
	"LogView/models"
	"github.com/astaxie/beego/logs"
	"fmt"
	"strings"
	"io"
	"strconv"
	"net/http"
	"bytes"
	"regexp"
	"bufio"
)



type InspectionController struct {
	BaseController
}
func (c *InspectionController) Prepare() {
	c.BaseController.Prepare()
	c.checkLogin()
}


//添加
func (c *InspectionController) Add()  {
	var Perject models.Perject
	dataList1,err :=Perject.GetAll()
	if err == nil {
		c.Data["PList"] = dataList1
	}
	var Env models.Env
	dataList2,err :=Env.GetAll()
	if err == nil {
		c.Data["EList"] = dataList2
	}
	c.TplName = "health/add.html"
}

//添加页面保存数据
func (c *InspectionController) Post() {
	var Health models.Health
	if err := c.ParseForm(&Health); err == nil {
		if id ,err :=Health.Add(&Health); err == nil {
			c.Data["json"] = id
		} else {
			c.Data["json"] = err
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

//返回全部数据
func (c *InspectionController) List() {
	Name :=c.GetString("Name")
	if Name !="" {
		var Health models.Health
		dataList,err :=Health.GetPerjectHealthListByName(Name)
		if err == nil {
			c.Data["List"] = dataList
		}
		logs.Info("dataList :", dataList)
		c.TplName = "health/list.html"
	}else {
		var Health models.Health
		dataList, err := Health.GetAll()
		if err == nil {
			c.Data["List"] = dataList
		}
		logs.Info("dataList :", dataList)
		c.TplName = "health/list.html"
	}
}


//修改
func (c *InspectionController) Edit() {
	id, _ := c.GetInt("Id", 0)
	var Health models.Health
	health, err := Health.GetHealthById(id)
	if err == nil {
		c.Data["Health"] = health
	}
	var Perject models.Perject
	dataList1,err :=Perject.GetAll()
	if err == nil {
		c.Data["PList"] = dataList1
	}
	var Env models.Env
	dataList2,err :=Env.GetAll()
	if err == nil {
		c.Data["EList"] = dataList2
	}
	c.TplName = "health/edit.html"
}

//保存
func (c *InspectionController) Update() {
	//自动解析绑定到对象中
	var Health models.Health
	if err := c.ParseForm(&Health); err == nil {
		if err :=Health.Update(&Health); err == nil {
			c.Data["json"] = ""
		} else {
			c.Data["json"] = "error"
		}
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}
//删除
func (c *InspectionController) Delete() {
	//获得id
	var Health models.Health
	id, _ := c.GetInt("Id", 0)
	if err := Health.Delete(id); err == nil {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

//巡检页面
func (c *InspectionController) Get()  {
	var Perject models.Perject
	dataList1,err :=Perject.GetAll()
	if err == nil {
		c.Data["PList"] = dataList1
	}
	var Env models.Env
	dataList2,err :=Env.GetAll()
	if err == nil {
		c.Data["EList"] = dataList2
	}
	c.TplName = "health.html"
}

//巡检方法函数
func (c *InspectionController) GetHealth()  {
	var Remote models.Remote
	var LogsContent bytes.Buffer
	PerjectId,PerErr :=c.GetInt("PerjectId")
	if PerErr != nil{
		logs.Error("项目ID:",PerjectId)
		Msg :="项目ID参数错误！"
		LogsContent.WriteString(fmt.Sprintf("%s:%s", Msg, PerjectId))
		c.Data["json"] = "{\"Message\":\"" + Msg + "\"}"
	}
	EnvId,EnvErr :=c.GetInt("EnvId")
	if EnvErr != nil {
		logs.Error("环境ID:",EnvId)
		Msg :="环境ID参数错误！"
		LogsContent.WriteString(fmt.Sprintf("%s:%s", Msg, EnvId))
		c.Data["json"] = "{\"Message\":\"" + Msg + "\"}"
	}
	var Health models.Health
	list,err :=Health.GetHealthByPerjectIdAndEnvId(PerjectId,EnvId)
	if err == nil {
		for _, v := range list {
			if v.Inspection == "http" {
				process_name, Err := Remote.GetSshClientByExec(v.User, v.Passwd, v.Machine, v.Hostport, "ps -eaf|grep -v grep |grep "+v.Processname+"|wc -l")
				if Err != nil {
					Msg := "进程检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------失败!\n"
					LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
				} else {
					ProcessNub, _ := strconv.Atoi(process_name)
					if ProcessNub > 0 {
						Msg := "进程检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------正常!\n"
						LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
					} else {
						Msg := "进程检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------异常!\n"
						LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
					}
				}
				application_port, Err := Remote.GetSshClientByExec(v.User, v.Passwd, v.Machine, v.Hostport, "netstat -tunlp|grep -v grep |grep "+strconv.Itoa(v.Port)+" |wc -l")
				if Err != nil {
					Msg := "端口检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------失败!\n"
					LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
				} else {
					ProcessNub, _ := strconv.Atoi(application_port)
					if ProcessNub > 0 {
						Msg := "端口检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------正常!\n"
						LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
					} else {
						Msg := "端口检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------异常!\n"
						LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
					}
				}
				res, err := http.Get("" + v.Url + "")
				if err != nil {
					Msg := "URL检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------失败!\n"
					LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
				}
				resCode := res.StatusCode
				defer res.Body.Close()
				if resCode == 200 {
					Msg := "URL检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------正常!\n"
					LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
				} else {
					Msg := "URL检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------异常!\n"
					LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
				}
				//disk检查
				Disk_information, Err := Remote.GetSshClientByExec(v.User, v.Passwd, v.Machine, v.Hostport, "df -lh|grep -v grep |grep -v 'tmpfs'|grep -v 'Filesystem'")
                if Err != nil {
					Msg := "磁盘检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------失败!\n"
					LogsContent.WriteString(fmt.Sprintf("%s--------------------------------%s","服务器IP："+v.Machine+"",Msg))
				}else {
					data :=strings.NewReader(Disk_information)
					reader := bufio.NewReader(data)
					var index int
					for {
						line, err2 :=reader.ReadString('\n')
						if err2 != nil || io.EOF == err2 {
							break
						}
						if 0 == len(line) || line == "\r\n"{
							continue
						}
						index++
						line = strings.Replace(line, "\n", "", -1)
						line =delete_extra_space(line)
						disk := strings.Split(line, " ")
						Msg :="检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------"+disk[4]+"\n"
						LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"","分区："+disk[5]+"",Msg))
					}
				}
			} else {
				process_name,Err:=Remote.GetSshClientByCMD(v.User,v.Passwd,v.Machine,v.Hostport,"ps -eaf|grep -v grep |grep "+v.Processname+"|wc -l")
				if Err !=nil {
					Msg :="进程检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------失败!\n"
					LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
				}else {
					ProcessNub,_ :=strconv.Atoi(strings.Replace(process_name, "\n", "", -1))
					if ProcessNub > 0 {
						Msg :="进程检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------正常!\n"
						LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
					}else {
						Msg :="进程检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------异常!\n"
						LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
					}
				}
				application_port,Err:=Remote.GetSshClientByCMD(v.User,v.Passwd,v.Machine,v.Hostport,"netstat -tunlp|grep -v grep |grep "+ strconv.Itoa(v.Port)+" |wc -l")
				if Err !=nil {
					Msg :="端口检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------失败!\n"
					LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
				}else {
					appNub,_ :=strconv.Atoi(strings.Replace(application_port, "\n", "", -1))
					if appNub > 0 {
						Msg :="端口检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------正常!\n"
						LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
					}else {
						Msg :="端口检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------异常!\n"
						LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"",v.Name,Msg))
					}
				}
				//disk检查
				Disk_information, Err := Remote.GetSshClientByExec(v.User, v.Passwd, v.Machine, v.Hostport, "df -lh|grep -v grep |grep -v 'tmpfs'|grep -v 'Filesystem'")
				if Err != nil {
					Msg := "磁盘检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------失败!\n"
					LogsContent.WriteString(fmt.Sprintf("%s--------------------------------%s","服务器IP："+v.Machine+"",Msg))
				}else {
					data :=strings.NewReader(Disk_information)
					reader := bufio.NewReader(data)
					var index int
					for {
						line, err2 :=reader.ReadString('\n')
						if err2 != nil || io.EOF == err2 {
							break
						}
						if 0 == len(line) || line == "\r\n"{
							continue
						}
						index++
						line = strings.Replace(line, "\n", "", -1)
						line =delete_extra_space(line)
						disk := strings.Split(line, " ")
						Msg :="检查--------------------------------------------------------------------------------------------------------------------------------------------------------------------"+disk[4]+"\n"
						LogsContent.WriteString(fmt.Sprintf("%s----------%s--------------------%s","服务器IP："+v.Machine+"","分区："+disk[5]+"",Msg))
					}
				}
			}
		}
	}
	var index int
	for {
		line, err2 := LogsContent.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		if 0 == len(line) || line == "\r\n" {
			continue
		}
		index++
		line = strings.Replace(line, "\n", "", -1)
		//fmt.Println("123",line)
		msg := models.Message{Message:line}
		broadcast <- msg
	}
	c.Data["json"] = "成功！"
	c.ServeJSON()

}
//删除多余空格的函数
func delete_extra_space(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "	", " ", -1)    //替换tab为空格
	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 { //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}