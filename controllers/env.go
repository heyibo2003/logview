package controllers

import (
	"LogView/models"
	"github.com/astaxie/beego/logs"
)

type EnvController struct {
	BaseController
}

func (c *EnvController) Prepare() {
	c.BaseController.Prepare()
	c.checkLogin()
}
func (c *EnvController) Add()  {
	c.TplName = "env/add.html"
}
//返回全部数据
func (c *EnvController) List() {
	var Env models.Env
	dataList, err :=Env.GetAll()
	if err == nil {
		c.Data["List"] = dataList
	}
	logs.Info("dataList :", dataList)
	c.TplName = "env/list.html"
}


//修改
func (c *EnvController) Edit() {
	id, _ := c.GetInt("Id", 0)
	var Env models.Env
	env, err := Env.GetEnvById(id)
	if err == nil {
		c.Data["Env"] = env
	}
	c.TplName = "env/edit.html"
}
//删除
func (c *EnvController) Delete() {
	//获得id
	var Env models.Env
	id, _ := c.GetInt("Id", 0)
	if err := Env.Delete(id); err == nil {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}


//添加页面保存数据
func (c *EnvController) Post() {
	var Env models.Env
	if err := c.ParseForm(&Env); err == nil {
		if id ,err :=Env.Add(&Env); err == nil {
			c.Data["json"] = id
		} else {
			c.Data["json"] = err
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

func (c *EnvController) Update() {
	//自动解析绑定到对象中
	var Env models.Env
	if err := c.ParseForm(&Env); err == nil {
		if err :=Env.Update(&Env); err == nil {
			c.Data["json"] = ""
		} else {
			c.Data["json"] = "error"
		}
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}