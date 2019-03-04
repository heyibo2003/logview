package controllers

import (
	"LogView/models"
	"github.com/astaxie/beego/logs"
)

type PerjectController struct {
	BaseController
}
func (c *PerjectController) Prepare() {
	c.BaseController.Prepare()
	c.checkLogin()
}

func (c *PerjectController) Add()  {
	c.TplName = "perject/add.html"
}

//返回全部数据
func (c *PerjectController) List() {
	var Perject models.Perject
	dataList, err :=Perject.GetAll()
	if err == nil {
		c.Data["List"] = dataList
	}
	logs.Info("dataList :", dataList)
	c.TplName = "perject/list.html"
}


//修改
func (c *PerjectController) Edit() {
	id, _ := c.GetInt("Id", 0)
	var Perject models.Perject
	perject, err := Perject.GetPerjectById(id)
	if err == nil {
		c.Data["Perject"] = perject
	}
	c.TplName = "perject/edit.html"
}

//删除
func (c *PerjectController) Delete() {
	//获得id
	var Perject models.Perject
	id, _ := c.GetInt("Id", 0)
	if err := Perject.Delete(id); err == nil {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

//添加页面保存数据
func (c *PerjectController) Post() {
	var Perject models.Perject
	if err := c.ParseForm(&Perject); err == nil {
		if id ,err :=Perject.Add(&Perject); err == nil {
			c.Data["json"] = id
		} else {
			c.Data["json"] = err
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

func (c *PerjectController) Update() {
	//自动解析绑定到对象中
	var Perject models.Perject
	if err := c.ParseForm(&Perject); err == nil {
		if err :=Perject.Update(&Perject); err == nil {
			c.Data["json"] = ""
		} else {
			c.Data["json"] = "error"
		}
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}