package controllers

import (
	"LogView/models"
	"github.com/astaxie/beego/logs"
)



type UserController struct {
	BaseController
}
func (c *UserController) Prepare() {
	c.BaseController.Prepare()
	c.checkLogin()
}
//添加
func (c *UserController) Add()  {
	var Perject models.Perject
	dataList,err :=Perject.GetAll()
	if err == nil {
		c.Data["List"] = dataList
	}
	logs.Info("dataList :", dataList)
	c.TplName = "user/add.html"
}

//添加页面保存数据
func (c *UserController) Post() {
	var User models.User
	if err := c.ParseForm(&User); err == nil {
		if id ,err :=User.AddOne(&User); err == nil {
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
func (c *UserController) List() {
	Name :=c.GetString("Name")
	if Name !="" {
		var User models.User
		dataList,err :=User.GetUserByName(Name)
		if err == nil {
			c.Data["List"] = dataList
		}
		logs.Info("dataList :", dataList)
		c.TplName = "user/list.html"
	}else {
		var User models.User
		dataList, err := User.GetAll()
		if err == nil {
			c.Data["List"] = dataList
		}
		logs.Info("dataList :", dataList)
		c.TplName = "user/list.html"
	}
}


//修改
func (c *UserController) Edit() {
	id, _ := c.GetInt("Id", 0)
	var User models.User
	user, err := User.GetUserById(id)
	if err == nil {
		c.Data["User"] = user
	}else {
		tmpUser :=models.User{}
		tmpUser.Status = -1
		tmpUser.Gender = -1
		c.Data["User"] = tmpUser
	}
	var Perject models.Perject
	dataList,err :=Perject.GetAll()
	if err == nil {
		c.Data["List"] = dataList
	}
	c.TplName = "user/edit.html"
}

//保存
func (c *UserController) Update() {
	//自动解析绑定到对象中
	var User models.User
	if err := c.ParseForm(&User); err == nil {
		if err :=User.Update(&User); err == nil {
			c.Data["json"] = "成功"
		} else {
			c.Data["json"] = "error"
		}
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}
//删除
func (c *UserController) Delete() {
	//获得id
	var User models.User
	id, _ := c.GetInt("Id", 0)
	if err := User.Delete(id); err == nil {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}


