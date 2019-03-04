package controllers

import (
	"LogView/models"
	"github.com/astaxie/beego/logs"
)



type SoftwareController struct {
	BaseController
}
func (c *SoftwareController) Prepare() {
	c.BaseController.Prepare()
	c.checkLogin()
}
//添加
func (c *SoftwareController) Add()  {
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
	c.TplName = "software/add.html"
}

//添加页面保存数据
func (c *SoftwareController) Post() {
	var Software models.Software
	if err := c.ParseForm(&Software); err == nil {
		if id ,err :=Software.AddOne(&Software); err == nil {
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
func (c *SoftwareController) List() {
	Domain :=c.GetString("Domain")
	if Domain !="" {
		var Software models.Software
		dataList,err :=Software.GetSoftwareByDomain(Domain)
		if err == nil {
			c.Data["List"] = dataList
		}
		logs.Info("dataList :", dataList)
		c.TplName = "software/list.html"
	}else {
		var Software models.Software
		dataList, err := Software.GetAll()
		if err == nil {
			c.Data["List"] = dataList
		}
		logs.Info("dataList :", dataList)
		c.TplName = "software/list.html"
	}
}


//修改
func (c *SoftwareController) Edit() {
	id, _ := c.GetInt("Id", 0)
	var Software models.Software
	software, err := Software.GetSoftwareById(id)
	if err == nil {
		c.Data["Software"] = software
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
	c.TplName = "software/edit.html"
}

//保存
func (c *SoftwareController) Update() {
	//自动解析绑定到对象中
	var Software models.Software
	if err := c.ParseForm(&Software); err == nil {
		if err :=Software.Update(&Software); err == nil {
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
func (c *SoftwareController) Delete() {
	//获得id
	var Software models.Software
	id, _ := c.GetInt("Id", 0)
	if err := Software.Delete(id); err == nil {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}







