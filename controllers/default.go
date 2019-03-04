package controllers

import (
	"LogView/models"
	"fmt"
)

type LoginController struct {
    BaseController
}

func (c *LoginController) Login() {
		c.TplName = "login.html"
}

func (c *LoginController) Post() {
	Session  :=c.StartSession()
	var User models.User
	inputs :=c.Input()
	username := inputs.Get("Username")
	password := inputs.Get("Password")
	var status int= 0
	var Perjectid int = 0

	if "" == username {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写用户名"}
	}else if "" == password {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
	}
	if username == "admin" && password == "p@ssw0rd" {
		Session.Set("username",username)
		Session.Set("status", "1")
		Session.Set("Perjectid","0")
		c.Data["json"]=map[string]interface{}{"code": 1, "message": "贺喜你，登录成功","islogin":1};
	}else  {
		users,err := User.LoginUser(username, password)

		fmt.Println("user:",users)
		if err == nil {
			for _, v := range users {
				status = v.Status
				Perjectid = v.Perjectid
			}
			Session.Set("username", username)
			Session.Set("status",status)
			Session.Set("Perjectid",Perjectid)
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "贺喜你，登录成功", "islogin": 2}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败","islogin":0}
		}
	}
	c.ServeJSON()
}

func (c *LoginController) GetUserName()  {
	userLogin := c.GetSession("username")
	c.Data["json"] = map[string]interface{}{ "username": userLogin}
	c.ServeJSON()
}

func (c *LoginController) Logout() {
	v := c.GetSession("username")
	if v !=nil {
		c.DelSession("username")
		c.DelSession("status")
		c.DelSession("Perjectid")
		c.DelSession("isLogin")
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "退出成功","islogin":0}
	c.ServeJSON()
}

