package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	isLogin int
}

func (c *BaseController) Prepare() {
     c.adapterUserInfo()
}


//从session里取用户信息
func (c *BaseController) adapterUserInfo() {
	userLogin := c.GetSession("username")
	Session  :=c.StartSession()
	if userLogin == nil {
		Session.Set("isLogin",0)
	} else if userLogin == "admin"{
		Session.Set("isLogin",1)
	}else {
		Session.Set("isLogin",2)
	}
}

func (c *BaseController) checkLogin() {
	isLogin :=c.GetSession("isLogin")
	if isLogin == 0  {
		//登录页面地址
		c.pageLogin()
	}
}


// 重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

// 重定向 去错误页
func (c *BaseController) pageError(msg string) {
	errorurl := c.URLFor("BaseController.Error") + "/" + msg
	c.Redirect(errorurl, 302)
	c.StopRun()
}

// 重定向 去登录页
func (c *BaseController) pageLogin() {
	url := c.URLFor("LoginController.Login")
	c.Redirect(url, 302)
	c.StopRun()
}

func (c *BaseController) Page404() {
	c.TplName = "page-404.html"
}

func (c *BaseController) Error() {
	c.Data["error"] = c.GetString(":error")
	c.TplName = "error.html"
}