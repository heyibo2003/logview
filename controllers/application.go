package controllers

import (
	"LogView/models"
	"github.com/astaxie/beego/logs"
)

type ApplicationController struct {
	BaseController
}
func (c *ApplicationController) Prepare() {
	c.BaseController.Prepare()
	c.checkLogin()
}
func (c *ApplicationController) Add()  {
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
	c.TplName = "application/add.html"
}
//返回全部数据
func (c *ApplicationController) List() {
	Domain :=c.GetString("Domain")
	if Domain !="" {
		var Environmentals models.Environmentals
		dataList, err := Environmentals.GetEnvironmentalsByDomain(Domain)
		if err == nil {
			c.Data["List"] = dataList
		}
		logs.Info("dataList :", dataList)
		c.TplName = "application/list.html"
	}else {
		var Environmentals models.Environmentals
		dataList, err := Environmentals.GetAll()
		if err == nil {
			c.Data["List"] = dataList
		}
		//logs.Info("dataList :", dataList)
		//for _, v := range dataList {
		//	logs.Info("dataList :", v)
		//	//dataList = append(dataList, v)
		//}
		c.TplName = "application/list.html"
	}
}


//修改
func (c *ApplicationController) Edit() {
	id, _ := c.GetInt("Id", 0)
	var Environmentals models.Environmentals
	environmentals, err := Environmentals.GetEnvironmentalsById(id)
	if err == nil {
		c.Data["Environmentals"] = environmentals
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
	c.TplName = "application/edit.html"
}
//删除
func (c *ApplicationController) Delete() {
	//获得id
	var Environmentals models.Environmentals
	id, _ := c.GetInt("Id", 0)
	if err := Environmentals.Delete(id); err == nil {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}


//添加页面保存数据
func (c *ApplicationController) Post() {
	var  Environmentals models.Environmentals
	if err := c.ParseForm(&Environmentals); err == nil {
		if id ,err :=Environmentals.Add(&Environmentals); err == nil {
			c.Data["json"] = id
		} else {
			c.Data["json"] = err
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

func (c *ApplicationController) Update() {
	//自动解析绑定到对象中
	var Environmentals models.Environmentals
	if err := c.ParseForm(&Environmentals); err == nil {
		if err :=Environmentals.Update(&Environmentals); err == nil {
			c.Data["json"] = ""
		} else {
			c.Data["json"] = "error"
		}
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}